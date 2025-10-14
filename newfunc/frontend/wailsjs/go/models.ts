export namespace main {
	
	export class usergovnas {
	    id: number;
	    name: string;
	    sex: string;
	    sumgavna: number;
	
	    static createFrom(source: any = {}) {
	        return new usergovnas(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.sex = source["sex"];
	        this.sumgavna = source["sumgavna"];
	    }
	}

}

