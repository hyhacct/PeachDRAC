export namespace model {
	
	export class ActionRequest {
	    action: string;
	    ips: string[];
	    // Go type: struct { Speed int "json:\"speed\"" }
	    fan: any;
	    // Go type: struct { Mount struct { IP string "json:\"ip\""; Path string "json:\"path\"" } "json:\"mount\"" }
	    nfs: any;
	
	    static createFrom(source: any = {}) {
	        return new ActionRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.action = source["action"];
	        this.ips = source["ips"];
	        this.fan = this.convertValues(source["fan"], Object);
	        this.nfs = this.convertValues(source["nfs"], Object);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

