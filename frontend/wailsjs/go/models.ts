export namespace main {
	
	export class ObfuscationConfig {
	    RemoveWhiteSpace: boolean;
	    RemoveComments: boolean;
	    ChangeVariable: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ObfuscationConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.RemoveWhiteSpace = source["RemoveWhiteSpace"];
	        this.RemoveComments = source["RemoveComments"];
	        this.ChangeVariable = source["ChangeVariable"];
	    }
	}

}

export namespace obfuscator {
	
	export class Config {
	    single_line_output: boolean;
	    string_literal: boolean;
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
	        this.single_line_output = source["single_line_output"];
	        this.string_literal = source["string_literal"];
	        this.loop_statement = source["loop_statement"];
	        this.if_statement = source["if_statement"];
	        this.constant_name = source["constant_name"];
	        this.variable_name = source["variable_name"];
	        this.function_name = source["function_name"];
	        this.remove_comments = source["remove_comments"];
	    }
	}

}

