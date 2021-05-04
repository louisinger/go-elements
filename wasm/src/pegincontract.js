const lib = require("./wallycore");

console.log(lib.preInit());
// lib.run();

// const h2b = (str) => Buffer.from(str, "hex");
// const fed_script = h2b(
//   "52210307fd375ed7cced0f50723e3e1a97bbe7ccff7318c815df4e99a59bc94dbcd819210367c4f666f18279009c941e57fab3e42653c6553e5ca092c104d1db279e328a2852ae"
// );

// const script = h2b("0014879008279c4e17fe0c61f9a84d82216cb81ddaff");
// const contract = lib._malloc(512);
// const res = lib.ccall(
//   "wally_elements_pegin_contract_script_from_bytes",
//   ["number"],
//   [
//     "array",
//     "number",
//     "array",
//     "number",
//     "number",
//     "number",
//     "number",
//     "number",
//   ],
//   [fed_script, fed_script.length, script, script.length, 0, contract, 512, 4]
// );
// console.log(res);
