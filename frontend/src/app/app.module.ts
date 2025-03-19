import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './pages/login/login.component';
import { SignupComponent } from './pages/signup/signup.component';
import { HttpClientModule } from '@angular/common/http';
import { ProfileComponent } from './pages/profile/profile.component';
import { SidebarComponent } from './pages/sidebar/sidebar.component';
import { EducationComponent } from './pages/education/education.component';
import { TeachingComponent } from './pages/teaching/teaching.component';
import { AnnouncementComponent } from './pages/class/announcement/announcement.component';
import { TopbarComponent } from './pages/class/topbar/topbar.component';
import { UsersComponent } from './pages/class/users/users.component';
import { TableComponent } from './pages/class/table/table.component';

@NgModule({
  declarations: [AppComponent, LoginComponent, SignupComponent, ProfileComponent, SidebarComponent, EducationComponent, TeachingComponent, AnnouncementComponent, TopbarComponent, UsersComponent, TableComponent],
  imports: [BrowserModule, AppRoutingModule, FormsModule, HttpClientModule, ReactiveFormsModule],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
