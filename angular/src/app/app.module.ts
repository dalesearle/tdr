import { NgModule } from '@angular/core';
import { AppRoutingModule } from './app-routing.module'
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule }    from '@angular/forms';
import { HttpModule} from '@angular/http';
import { RouterModule } from '@angular/router';


import { AppComponent } from './app.component';
import { AppErrorComponent } from "./app-error.component";
import { CreateAccountComponent } from "./create-account.component"
import { LoginComponent } from "./login.component";
import {MenuComponent} from "./menu.component";


@NgModule({
  declarations: [
    AppComponent,
    AppErrorComponent,
    CreateAccountComponent,
    LoginComponent,
    MenuComponent,
  ],
  imports: [
    AppRoutingModule,
    BrowserModule,
    FormsModule,
    HttpModule,
    RouterModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})

export class AppModule { }
