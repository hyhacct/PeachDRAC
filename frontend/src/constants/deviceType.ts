const type = [
  {
    label: '戴尔',
    value: 'success'
  },
  {
    label: '浪潮',
    value: 'info'
  },
];

export const deviceType = (name: string) => {
  return type.find((item) => item.label === name)?.value || 'error';
};