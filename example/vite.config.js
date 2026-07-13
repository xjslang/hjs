import { defineConfig } from "vite";
import { execFileSync } from "child_process";

function hjsPlugin() {
  return {
    name: "vite-plugin-hjs",
    load(id) {
      const file = id.split("?", 1)[0];
      if (file.endsWith(".hjs")) {
        const compiled = execFileSync("hjs", [file], { encoding: "utf-8" });
        return {
          code: compiled,
          map: null,
        };
      }
    },
  };
}

export default defineConfig({
  plugins: [hjsPlugin()],
});
