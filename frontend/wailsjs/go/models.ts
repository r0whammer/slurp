export namespace spotify {
	
	export class SavedTrackRow {
	    id: string;
	    coverUrl: string;
	    trackTitle: string;
	    artistTitle: string;
	    albumTitle: string;
	    addedAt: string;
	
	    static createFrom(source: any = {}) {
	        return new SavedTrackRow(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.coverUrl = source["coverUrl"];
	        this.trackTitle = source["trackTitle"];
	        this.artistTitle = source["artistTitle"];
	        this.albumTitle = source["albumTitle"];
	        this.addedAt = source["addedAt"];
	    }
	}

}

