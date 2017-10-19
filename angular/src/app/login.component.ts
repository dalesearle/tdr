import { Component, Input} from '@angular/core';
import { Account } from './account';
import { AccountService } from './account.service';
import { Router } from '@angular/router';

@Component({
  selector: 'login',
  template: `
    <div>
      <form action="/login" #loginForm = "ngForm" (ngSubmit)="onSubmit()">
        <fieldset>
          <legend align="center">Login</legend>
          <label class="err_msg">{{ error }}</label><br>
          <div>
            <label>Email:</label>
            <input type="email" [(ngModel)]="acct.email" name="email" email autocomplete="on" autofocus="on" required>
          </div>
          <div>
            <label>Password:</label>
            <input type="password" [(ngModel)]="acct.pwd" name="pwd" autocomplete="on" pattern=".{6,}" required>
          </div>
          <div>
            <button type="submit" [disabled]="!loginForm.form.valid">Submit</button>
          </div>
          <a href="/create-account">Create Account</a>
        </fieldset>
      </form>
    </div>`,
  styleUrls: ['./login.component.css'],
  providers: [AccountService]
})

export class LoginComponent{
  @Input() acct: Account;
  error: string;

  constructor(
    private acctSvc: AccountService,
    private router: Router,
  ) {
    this.acct = new Account();
  }

  onSubmit(){
    this.acctSvc.login(this.acct).subscribe(
      resp => this.processResponse(resp.status),
      err => this.processError(err.status)
    );
  }

  processResponse(status: number){
    this.router.navigateByUrl('/menu')
  }

  processError(status: number){
    switch (status) {
      case 401:
      case 404:
        this.error = "Credentials Failed Authentication"
        break;
      case 500:
        this.router.navigateByUrl('/app-error')
        break;
    }
    this.acct.clear()
  }
}
