import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './pages/login/login.component';
import { SignupComponent } from './pages/signup/signup.component';
import { ProfileComponent } from './pages/profile/profile.component';
import { EducationComponent } from './pages/education/education.component';
import { TeachingComponent } from './pages/teaching/teaching.component';
import { AnnouncementComponent } from './pages/class/announcement/announcement.component';
import { UsersComponent } from './pages/class/users/users.component';
import { TableComponent } from './pages/class/table/table.component';
import { AssignmentsListComponent } from './pages/class/assignments-list/assignments-list.component';
import { AssignmentComponent } from './pages/class/assignment/assignment.component';
import { PerformanceComponent } from './pages/performance/performance.component';

import { AuthGuard } from './auth.guard';

const routes: Routes = [
  { path: 'log-in', component: LoginComponent },
  { path: 'sign-up', component: SignupComponent },
  { path: '', component: EducationComponent, canActivate: [AuthGuard] },
  { path: 'profile', component: ProfileComponent, canActivate: [AuthGuard] },
  { path: 'teaching', component: TeachingComponent, canActivate: [AuthGuard] },
  { path: 'performance', component: PerformanceComponent, canActivate: [AuthGuard] },
  { path: 'class/:id', component: AnnouncementComponent, canActivate: [AuthGuard] },
  { path: 'class/:id/users-list', component: UsersComponent, canActivate: [AuthGuard] },
  { path: 'class/:id/table', component: TableComponent, canActivate: [AuthGuard] },
  { path: 'class/:id/assignments', component: AssignmentsListComponent, canActivate: [AuthGuard] },
  { path: 'class/:id/assignment/:assignmentId', component: AssignmentComponent, canActivate: [AuthGuard] },
  // { path: '', component: MainComponent, canActivate: [AuthGuard]} ,
  // { path: '**', redirectTo: 'log-in'},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
