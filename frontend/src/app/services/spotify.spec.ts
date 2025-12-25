import { TestBed } from '@angular/core/testing';
import { SpotifyService } from './spotify';

declare const window: any;

describe('SpotifyService', () => {
  let service: SpotifyService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SpotifyService);
  });

  afterEach(() => {
    // clean up between tests
    window.go = undefined;
  });

  it('returns saved tracks from the wails binding', async () => {
    const mockRows = [
      {
        id: '1',
        coverUrl: 'https://example.com/a.jpg',
        trackTitle: 'Track A',
        artistTitle: 'Artist A',
        albumTitle: 'Album A',
        addedAt: '2025-12-24T00:00:00Z',
      },
    ];

    const spy = jasmine.createSpy('GetSavedTracks').and.resolveTo(mockRows);
    window.go = { main: { App: { GetSavedTracks: spy } } };

    const rows = await service.getSavedTracks(50, 0);

    expect(spy).toHaveBeenCalledOnceWith(50, 0);
    expect(rows.length).toBe(1);
    expect(rows[0].trackTitle).toBe('Track A');
  });

  it('throws if the wails binding is missing', async () => {
    window.go = {}; // no binding exposed

    await expectAsync(service.getSavedTracks(10, 0)).toBeRejectedWithError(
      /GetSavedTracks not found/i
    );
  });
});
