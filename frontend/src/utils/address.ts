
/*
  验证IP是否合法
  @param ip: string
  @return: boolean
*/
function verify_ip(ip: string) {
  const ip_regex = /^(\d{1,3}\.){3}\d{1,3}$/;
  return ip_regex.test(ip);
}


/*
  生成IP列表,根据IP地址+网段生成1-255的IP列表,如果存在错误则会抛出错误
  @param ip: string
  @param paragraph: string
  @return: string[]
*/
function generate_ip_list(ip: string, paragraph: string) {
  const ip_list = [];
  const ip_parts = ip.split('.');
  if (ip_parts.length !== 4) {
    throw new Error('IP地址格式不正确');
  }
  for (let i = 1; i < 256; i++) {
    switch (paragraph) {
      case '1':
        ip_list.push(`${i}.${ip_parts[1]}.${ip_parts[2]}.${ip_parts[3]}`);
        break;
      case '2':
        ip_list.push(`${ip_parts[0]}.${i}.${ip_parts[2]}.${ip_parts[3]}`);
        break;
      case '3':
        ip_list.push(`${ip_parts[0]}.${ip_parts[1]}.${i}.${ip_parts[3]}`);
        break;
      case '4':
        ip_list.push(`${ip_parts[0]}.${ip_parts[1]}.${ip_parts[2]}.${i}`);
        break;
      default:
        throw new Error('网段格式不正确');
    }
  }
  return ip_list;
}


export {
  verify_ip,
  generate_ip_list,
};
