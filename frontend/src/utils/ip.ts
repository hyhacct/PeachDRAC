class IPTraverser {
    // 将 IP 字符串分解为数组
    private parseIp(ip: string): number[] {
        return ip.split('.').map(Number);
    }

    // 验证 IP 是否有效
    private isValidIp(ip: string): boolean {
        const parts = ip.split('.');
        if (parts.length !== 4) return false;
        
        return parts.every(part => {
            const num = Number(part);
            return !isNaN(num) && num >= 0 && num <= 255 && String(num) === part;
        });
    }

    // 生成 IP 字符串
    private buildIp(parts: number[]): string {
        return parts.join('.');
    }

    // 遍历指定网段
    traverse(startIp: string, changingOctet: number): string[] {
        // 验证输入
        if (!this.isValidIp(startIp)) {
            throw new Error('Invalid IP address format');
        }
        if (changingOctet < 0 || changingOctet > 3) {
            throw new Error('Changing octet must be between 0 and 3');
        }

        const startParts = this.parseIp(startIp);
        const result: string[] = [];
        const baseIp = [...startParts];

        // 从起始值到 254 遍历指定段
        const startValue = startParts[changingOctet];
        for (let value = startValue; value <= 254; value++) {
            baseIp[changingOctet] = value;
            result.push(this.buildIp(baseIp));
        }

        return result;
    }
}

export default IPTraverser;