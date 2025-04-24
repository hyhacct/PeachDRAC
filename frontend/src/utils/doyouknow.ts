const text = [
  "启用的密码组越多,探测速度越慢,这是因为探测机器时会将所有密码全尝试一遍,直到登录成功~",
  "如果你觉得等待太久了,可以切换到其他页面做其他操作,例如批量动作,探测状态不会被打断~",
  "探测结果有时候并不准确,受到网络的影响,可以尝试多进行几次探测~",
]


// 随机取一段话出来
export function getRandomText() {
  return text[Math.floor(Math.random() * text.length)]
}
