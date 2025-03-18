import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './pages/login/login.component';
import { SignupComponent } from './pages/signup/signup.component';
import { ProfileComponent } from './pages/profile/profile.component';
import { EducationComponent } from './pages/education/education.component';
import { TeachingComponent } from './pages/teaching/teaching.component';
import { AnnouncementComponent } from './pages/class/announcement/announcement.component';

import { AuthGuard } from './auth.guard';

const routes: Routes = [
  { path: 'log-in', component: LoginComponent },
  { path: 'sign-up', component: SignupComponent },
  { path: '', component: EducationComponent, canActivate: [AuthGuard] },
  { path: 'profile', component: ProfileComponent, canActivate: [AuthGuard] },
  { path: 'teaching', component: TeachingComponent, canActivate: [AuthGuard] },
  { path: 'class/:id', component: AnnouncementComponent, canActivate: [AuthGuard] },
  // { path: '', component: MainComponent, canActivate: [AuthGuard]} ,
  // { path: '**', redirectTo: 'log-in'},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
