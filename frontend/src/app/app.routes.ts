import { Routes } from '@angular/router';
import { DashboardComponent } from './dashboard/dashboard';

export const routes: Routes = [
  { path: '', pathMatch: 'full', redirectTo: 'home' },
  { path: 'dashboard', component: DashboardComponent},

  { path: '**', redirectTo: 'dashboard' },
];
