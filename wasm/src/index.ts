import "./wasm_exec.cjs";
import { readFileSync } from "fs";

const WASM_URL = "src/main.wasm"; // TODO improve this??
// @ts-ignore
const go = new Go();

export interface PeginModule {
  getPeginAddress: (pubKey: string, fedPegScript: string, network: number, isDynaFed: boolean, contract: string) => { claimScript: string, mainChainAddress: string }
}

/**
 * return the WebAssembly.Instance according to current environment (web or node)
 */
async function webAssemblyInstance(): Promise<WebAssembly.Instance | WebAssembly.WebAssemblyInstantiatedSource> {
  if (typeof window === 'undefined') {
    // node
    const mod = await WebAssembly.compile(readFileSync(WASM_URL));
    return WebAssembly.instantiate(mod, go.importObject);
  }
  // web
  return WebAssembly.instantiateStreaming(fetch(WASM_URL), go.importObject)
}

/**
 * get the instance, run wasm and return PeginModule
 */
export async function loadModule(): Promise<PeginModule> {
  const instance = await webAssemblyInstance();
  go.run(instance);
  return {
    // @ts-ignore
    getPeginAddress: getPeginAddress
  }
}
