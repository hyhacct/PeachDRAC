export namespace model {

	export class WailsCommunicate {
		Status: boolean;
		Msg: string;
		Data: string;

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

