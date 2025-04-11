export const buttons = [
  {
    name: '打开控制台',
    icon: 'hugeicons:java',
    color: '#FF7F50',
    action: 'openConsole',
    isWarning: false, // 是否显示警告
  },
  {
    name: '开机',
    icon: 'material-symbols:power-settings-new',
    color: '#F5F5DC',
    action: 'powerOn',
    isWarning: false, // 是否显示警告
  },
  {
    name: '关机',
    icon: 'lucide:power-off',
    color: '#F5F5DC',
    action: 'powerOff',
    isWarning: true, // 是否显示警告
  },
  {
    name: '重启',
    icon: 'hugeicons:hydro-power',
    color: '#F5F5DC',
    action: 'powerReset',
    isWarning: true, // 是否显示警告
  },
  {
    name: '冷引导',
    icon: 'streamline:image-flash-1-flash-power-connect-charge-electricity-lightning',
    color: '#F5F5DC',
    action: 'powerColdBoot',
    isWarning: true, // 是否显示警告
  },
  {
    name: '风扇调整',
    icon: 'ix:device-fan',
    color: '#9370DB',
    action: 'fanAdjust',
    isWarning: false,
  },
  {
    name: '风扇自适应',
    icon: 'bi:fan',
    color: '#9370DB',
    action: 'fanAdaptive',
    isWarning: false, // 是否显示警告
  },
  {
    name: 'NFS挂载',
    icon: 'arcticons:nfs-no-limits',
    color: '#A9A9A9',
    action: 'mountNFS',
    isWarning: false, // 是否显示警告
  },
  {
    name: 'NFS卸载',
    icon: 'arcticons:nfs-no-limits',
    color: '#A9A9A9',
    action: 'unmountNFS',
    isWarning: false, // 是否显示警告
  },
];