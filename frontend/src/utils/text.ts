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


/**
 * 字符串截断，超过10个字符就截断
 * @param str 字符串
 * @param length 截断长度
 * @returns 截断后的字符串
 */
function cutOff(str: string, length: number = 10) {
  if (str?.length > length) {
    return str.slice(0, length) + '...';
  }
  return str;
}

export {
  isEmpty,
  splitIp,
  cutOff,
};

