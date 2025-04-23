export namespace model {
	
	export class WailsCommunicate {
	    status: boolean;
	    msg: string;
	    data: string;
	
	    static createFrom(source: any = {}) {
	        return new WailsCommunicate(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	        this.msg = source["msg"];
	        this.data = source["data"];
	    }
	}

}

