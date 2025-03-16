export namespace obfuscator {
	
	export class Config {
	    minify: boolean;
	    unicode_string: boolean;
	    loop_statement: boolean;
	    if_statement: boolean;
	    constant_name: boolean;
	    variable_name: boolean;
	    function_name: boolean;
	    remove_comments: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.minify = source["minify"];
	        this.unicode_string = source["unicode_string"];
	        this.loop_statement = source["loop_statement"];
	        this.if_statement = source["if_statement"];
	        this.constant_name = source["constant_name"];
	        this.variable_name = source["variable_name"];
	        this.function_name = source["function_name"];
	        this.remove_comments = source["remove_comments"];
	    }
	}

}

