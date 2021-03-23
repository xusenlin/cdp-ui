export const Phone = {
  pattern: /^1((3[\d])|(4[5,6,7,9])|(5[0-3,5-9])|(6[5-7])|(7[0-8])|(8[\d])|(9[1,8,9]))\d{8}$/,
  message: "请输入正确的手机号码",
  trigger: "blur"
};

export const Number = {
  type: "number",
  message: "只能填写数字",
  trigger: "blur"
};

export const String = {
  type: "string",
  message: "只能填写字符串",
  trigger: "blur"
};

export const Required = {
  required: true,
  message: "填写不能为空",
  trigger: "blur"
};

export const Array = {
  type: "array",
  message: "请选择选项",
  trigger: "change"
};

export const Date = { type: "date", message: "请选择日期", trigger: "change" };

export const Email = {
  type: "email",
  message: "请输入正确的邮箱",
  trigger: "blur"
};

export const StrLen3To18 = {
  pattern: /^.{3,18}$/,
  message: "字符长度限制在2到18个",
  trigger: "blur"
};

export const Float = {
  pattern: /^(([1-9][0-9]*)|(([0]\.\d{1,2}|[1-9][0-9]*\.\d{1,2})))$/,
  message: "只能填写数字并且最多两位小数",
  trigger: "blur"
};

export const FloatAnd0 = {
  pattern: /(^[1-9]([0-9]+)?(\.[0-9]{1,2})?$)|(^(0){1}$)|(^[0-9]\.[0-9]([0-9])?$)/,
  message: "只能填写数字并且最多两位小数",
  trigger: "blur"
};

export const FloatGeq1 = {
  pattern: /(^[1-9]([0-9]+)?(\.[0-9]{1,2})?$)|(^[1-9]\.[0-9]([0-9])?$)/,
  message: "只能填写数字并且大于等于1，最多两位小数",
  trigger: "blur"
};

export const FloatMax100 = {
  pattern: /^\d\.([1-9]{1,2}|[0-9][1-9])$|^[1-9]\d{0,1}(\.\d{1,2}){0,1}$|^100(\.0{1,2}){0,1}$/,
  message: "只能填写数字并且最大数是100不能超个两位小数",
  trigger: "blur"
};

export const Url = {
  pattern: /(^https?|^nurse):\/\//i,
  message: "请输入正确的链接地址",
  trigger: "blur"
};

export const IntMax99 = {
  pattern: /^[1-9][0-9]{0,1}$/,
  message: "只能填写1-99的整数",
  trigger: "blur"
};

export const FloatMin0Max100 = {
  pattern: /^(\d|[1-9]\d|100)(\.\d{1,2})?$/,
  message: "只能填写0-100，不能超个两位小数",
  trigger: "blur"
};

export const StrongPassword = {
  pattern: /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{6,18}$/,
  message: "大小写+数字组成，字数6-18位",
  trigger: "blur"
};

export const NumberString = {
  pattern: /^\d+$/,
  message: "只能输入数字",
  trigger: "blur"
};

export const Geq1NumberString = {
  pattern: /^[1-9][0-9]*$/,
  message: "只能填写大于等于1的整数",
  trigger: "blur"
};

export const PasswordLength = {
  pattern: /^[0-9a-zA-Z]{8,20}$/,
  message: "数字或字母组成，字数8-20位",
  trigger: "blur"
};

export function FillerFieldRules(fields = [], rules) {
  let newRules = {};
  fields.forEach(f => {
    newRules[f] = rules;
  });
  return newRules;
}
