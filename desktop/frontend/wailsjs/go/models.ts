export namespace api {
	
	export class ImageConfig {
	
	
	    static createFrom(source: any = {}) {
	        return new ImageConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

export namespace main {
	
	export class Conf {
	
	
	    static createFrom(source: any = {}) {
	        return new Conf(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

