/** 
 * 字符串是否为空
 * @param str 字符串
 * @returns 是否为空
 */
function isEmpty(str: string) {
  return str?.trim() === '';
}


/**
 * 分割IP地址
 * @param str 字符串
 * @returns 分割后的IP地址数组
 */
function splitIp(str: string) {
  return str?.split(',');
}

export {
  isEmpty,
  splitIp,
};
