const type = [
  {
    label: '在线',
    value: 'success'
  },
  {
    label: '离线',
    value: 'error'
  },
];

export const statusType = (name: string) => {
  return type.find((item) => item.label === name)?.value || '';
};