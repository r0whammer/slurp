import { CommonModule } from '@angular/common';
import { ChangeDetectorRef, Component, NgZone, OnInit, ViewChild } from '@angular/core';

import { Table, TableModule } from 'primeng/table';
import { InputTextModule } from 'primeng/inputtext';
import { ButtonModule } from 'primeng/button';
import { ProgressSpinnerModule } from 'primeng/progressspinner';

import { SpotifyService, SavedTrackRow } from '../services/spotify';

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [CommonModule, TableModule, InputTextModule, ButtonModule, ProgressSpinnerModule],
  templateUrl: './dashboard.html',
  styleUrls: ['./dashboard.scss'],
})
export class DashboardComponent implements OnInit {
  @ViewChild('dt') dt!: Table;

  tracks: SavedTrackRow[] = [];
  loading = true;
  error: string | null = null;

  readonly initialLimit = 200;

  constructor(
    private spotify: SpotifyService,
    private zone: NgZone,
    private cdr: ChangeDetectorRef,
  ) {}

  async ngOnInit(): Promise<void> {
    await this.refresh();
  }

  async refresh(): Promise<void> {
    this.zone.run(() => {
      this.loading = true;
      this.error = null;
    });

    try {
      const rows = await this.spotify.getSavedTracks(this.initialLimit, 0);

      this.zone.run(() => {
        this.tracks = rows;
        this.loading = false;
      });
    } catch (e: any) {
      this.zone.run(() => {
        this.error = e?.message ?? String(e);
        this.loading = false;
      });
    } finally {
      this.zone.run(() => {
        this.loading = false;
        this.cdr.markForCheck();
      });
    }
  }

  onGlobalFilter(event: Event): void {
    const input = event.target as HTMLInputElement;
    this.dt.filterGlobal(input.value, 'contains');
  }

  clearAll(): void {
    this.dt.clear();
  }
}
