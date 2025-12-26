export namespace model {
	
	export class HooksConfig {
	    Before: string[];
	    After: string[];
	
	    static createFrom(source: any = {}) {
	        return new HooksConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Before = source["Before"];
	        this.After = source["After"];
	    }
	}
	export class GenConfig {
	    Hooks?: HooksConfig;
	    Style: string;
	    GitChange: boolean;
	    Route2Code: boolean;
	    ModelDriver: string;
	    ModelCache: boolean;
	    ModelIgnoreCols: string[];
	    ModelSchema: string;
	
	    static createFrom(source: any = {}) {
	        return new GenConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Hooks = this.convertValues(source["Hooks"], HooksConfig);
	        this.Style = source["Style"];
	        this.GitChange = source["GitChange"];
	        this.Route2Code = source["Route2Code"];
	        this.ModelDriver = source["ModelDriver"];
	        this.ModelCache = source["ModelCache"];
	        this.ModelIgnoreCols = source["ModelIgnoreCols"];
	        this.ModelSchema = source["ModelSchema"];
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
	
	export class NewConfig {
	    Home: string;
	
	    static createFrom(source: any = {}) {
	        return new NewConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Home = source["Home"];
	    }
	}
	export class JZeroConfig {
	    Hooks?: HooksConfig;
	    New?: NewConfig;
	    Gen?: GenConfig;
	
	    static createFrom(source: any = {}) {
	        return new JZeroConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Hooks = this.convertValues(source["Hooks"], HooksConfig);
	        this.New = this.convertValues(source["New"], NewConfig);
	        this.Gen = this.convertValues(source["Gen"], GenConfig);
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

export namespace types {
	
	export class CommandResult {
	    success: boolean;
	    exitCode: number;
	    output: string;
	    error?: string;
	    duration: number;
	    // Go type: time
	    startTime: any;
	    // Go type: time
	    endTime: any;
	
	    static createFrom(source: any = {}) {
	        return new CommandResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.exitCode = source["exitCode"];
	        this.output = source["output"];
	        this.error = source["error"];
	        this.duration = source["duration"];
	        this.startTime = this.convertValues(source["startTime"], null);
	        this.endTime = this.convertValues(source["endTime"], null);
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
	export class FileInfo {
	    name: string;
	    path: string;
	    type: string;
	    size: number;
	    // Go type: time
	    modified: any;
	    ext: string;
	    isHidden: boolean;
	
	    static createFrom(source: any = {}) {
	        return new FileInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.path = source["path"];
	        this.type = source["type"];
	        this.size = source["size"];
	        this.modified = this.convertValues(source["modified"], null);
	        this.ext = source["ext"];
	        this.isHidden = source["isHidden"];
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
	export class FileNode {
	    id: string;
	    name: string;
	    path: string;
	    type: string;
	    ext?: string;
	    size?: number;
	    modified?: number;
	    children?: FileNode[];
	    isLeaf: boolean;
	
	    static createFrom(source: any = {}) {
	        return new FileNode(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.path = source["path"];
	        this.type = source["type"];
	        this.ext = source["ext"];
	        this.size = source["size"];
	        this.modified = source["modified"];
	        this.children = this.convertValues(source["children"], FileNode);
	        this.isLeaf = source["isLeaf"];
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
	export class ProjectInfo {
	    name: string;
	    path: string;
	    rootPath: string;
	    isPlugin: boolean;
	    parent?: string;
	    hasConfig: boolean;
	    descDirs: string[];
	    apiCount: number;
	    protoCount: number;
	    sqlCount: number;
	
	    static createFrom(source: any = {}) {
	        return new ProjectInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.path = source["path"];
	        this.rootPath = source["rootPath"];
	        this.isPlugin = source["isPlugin"];
	        this.parent = source["parent"];
	        this.hasConfig = source["hasConfig"];
	        this.descDirs = source["descDirs"];
	        this.apiCount = source["apiCount"];
	        this.protoCount = source["protoCount"];
	        this.sqlCount = source["sqlCount"];
	    }
	}
	export class ProjectNode {
	    id: string;
	    name: string;
	    path: string;
	    type: string;
	    isLeaf: boolean;
	    children?: ProjectNode[];
	    configPath?: string;
	    hasConfig: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ProjectNode(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.path = source["path"];
	        this.type = source["type"];
	        this.isLeaf = source["isLeaf"];
	        this.children = this.convertValues(source["children"], ProjectNode);
	        this.configPath = source["configPath"];
	        this.hasConfig = source["hasConfig"];
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

