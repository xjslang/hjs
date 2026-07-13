export const createContext = () => {
  const ctx = {
    currentEffect: {
      id: "",
      fn: undefined,
      effectsId: "",
    },
    cleaners: {},
  };

  return {
    signal: (val) => createSignal(ctx, val),
    effects: () => createEffects(ctx),
  };
};

const createSignal = (ctx, val) => {
  const signal = {
    id: genId(),
    val,
    effects: {},
  };

  function cleaner(effectsId) {
    for (const key in signal.effects) {
      if (effectsId === signal.effects[key].effectsId) {
        delete signal.effects[key];
      }
    }
  }

  return {
    get: () => {
      const { id, fn, effectsId } = ctx.currentEffect;
      signal.effects[id] = { fn, effectsId };
      ctx.cleaners[signal.id] = cleaner;
      return signal.val;
    },
    set: (val) => {
      signal.val = val;
      for (const key in signal.effects) {
        signal.effects[key].fn?.();
      }
    },
  };
};

const createEffects = (ctx) => {
  const effects = {
    id: genId(),
    signalCleaners: {},
  };

  return {
    effect: (fn) => {
      const id = genId();
      ctx.currentEffect = { id, fn, effectsId: effects.id };
      ctx.cleaners = effects.signalCleaners;
      fn();
    },
    clean: () => {
      for (const key in effects.signalCleaners) {
        effects.signalCleaners[key](effects.id);
      }
    },
  };
};

const genId = () => (+new Date()).toString();
