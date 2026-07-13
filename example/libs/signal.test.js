import { beforeEach, describe, expect, test } from "vitest";
import { createContext } from "./signal";

test("effects are called immediately", () => {
  const ctx = createContext();
  const msg = ctx.signal("hello there!");
  const efs1 = ctx.effects();
  const efs2 = ctx.effects();

  let out1 = "";
  efs1.effect(function () {
    out1 = msg.get();
  });

  let out2 = "";
  efs2.effect(function () {
    out2 = msg.get();
  });

  expect(out1).toBe("hello there!");
  expect(out2).toBe("hello there!");
});

test("changing a signal re-execute effects", () => {
  const ctx = createContext();
  const msg = ctx.signal("hello there!");
  const efs1 = ctx.effects();
  const efs2 = ctx.effects();

  // effects are called immediately
  let out1 = "";
  efs1.effect(function () {
    out1 = msg.get();
  });

  // effects are called immediately
  let out2 = "";
  efs2.effect(function () {
    out2 = msg.get();
  });

  expect(out1).toBe("hello there!");
  expect(out2).toBe("hello there!");
  // changing a signal re-execute effects
  msg.set("blah!");
  expect(out1).toBe("blah!");
  expect(out2).toBe("blah!");
});

test("after cleanup, the effects are no longer re-executed", () => {
  const ctx = createContext();
  const msg = ctx.signal("hello there!");
  const efs1 = ctx.effects();
  const efs2 = ctx.effects();

  let out1 = "";
  efs1.effect(function () {
    out1 = msg.get();
  });

  let out2 = "";
  efs2.effect(function () {
    out2 = msg.get();
  });

  expect(out1).toBe("hello there!");
  expect(out2).toBe("hello there!");

  // after cleanup, the effects are no longer re-executed
  efs1.clean();
  msg.set("aaa");
  expect(out1).toBe("hello there!"); // remains the prev value
  expect(out2).toBe("hello there!"); // remains the prev value
});
