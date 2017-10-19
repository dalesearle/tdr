import { NgModule }             from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import {AppErrorComponent} from "./app-error.component";
import { CreateAccountComponent } from './create-account.component';
import { LoginComponent } from './login.component';
import {MenuComponent } from './menu.component';

const routes: Routes = [
  { path: '', redirectTo: '/login', pathMatch: 'full' },
  { path: 'app-error', component: AppErrorComponent },
  { path: 'create-account', component: CreateAccountComponent },
  { path: 'login',  component: LoginComponent },
  {path: 'menu', component: MenuComponent}
]
@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})

export class AppRoutingModule {}
