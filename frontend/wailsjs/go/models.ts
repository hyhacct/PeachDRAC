export namespace model {
	
	export class TableJava {
	    id: number;
	    title: string;
	    path: string;
	    allot: string;
	    ips: string;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new TableJava(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.path = source["path"];
	        this.allot = source["allot"];
	        this.ips = source["ips"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}
	export class TablePass {
	    id: number;
	    username: string;
	    password: string;
	    port: number;
	    status: boolean;
	    priority: number;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new TablePass(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.username = source["username"];
	        this.password = source["password"];
	        this.port = source["port"];
	        this.status = source["status"];
	        this.priority = source["priority"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}
	export class WailsCommunicate {
	    Status: boolean;
	    Msg: string;
	    Data: any;
	
	    static createFrom(source: any = {}) {
	        return new WailsCommunicate(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Status = source["Status"];
	        this.Msg = source["Msg"];
	        this.Data = source["Data"];
	    }
	}

}

