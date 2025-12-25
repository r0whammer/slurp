import { Injectable } from '@angular/core';

// adjust this path if your bound struct is not package main / struct App
import { GetSavedTracks } from '../../../wailsjs/go/main/App';

export type SavedTrackRow = {
  id: string;
  coverUrl: string;
  trackTitle: string;
  artistTitle: string;
  albumTitle: string;
  addedAt?: string;
};

@Injectable({ providedIn: 'root' })
export class SpotifyService {
  async getSavedTracks(limit: number, offset: number): Promise<SavedTrackRow[]> {
    return (await GetSavedTracks(limit, offset)) as SavedTrackRow[];
  }
}

