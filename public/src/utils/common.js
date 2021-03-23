/**
 * 重置对象参数
 * @params -> Object
 * @arg = Array => []
 * @arg = Boolean => false
 * @arg = Number => null
 * @arg = String => ''
 * */
export function resetArgs(args, def = {}) {
  for (let key in args) {
    console.log(key);
    if (def.hasOwnProperty(key)) {
      args[key] = def[key];
    } else {
      if (Array.isArray(args[key])) {
        args[key] = [];
      }
      if ("string" == typeof args[key]) {
        args[key] = "";
      }
      if ("number" == typeof args[key]) {
        args[key] = null;
      }
      if ("boolean" == typeof args[key]) {
        args[key] = false;
      }
      if (Object.prototype.toString.call(args[key]) === "[object Object]") {
        args[key] = {};
      }
    }
  }
}

/**
 * 填充对象属性
 * @param obj
 * @param row
 * row属性赋值到obj
 */
export function fillerLeft(obj, row = {}) {
  for (let key in obj) {
    if (
      row.hasOwnProperty(key) &&
      row[key] !== null &&
      row[key] !== undefined
    ) {
      obj[key] = row[key];
    }
  }
}
export function clone(obj) {
  return JSON.parse(JSON.stringify(obj));
}
