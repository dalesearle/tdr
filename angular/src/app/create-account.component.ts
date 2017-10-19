import {Component, Input} from '@angular/core';
import {Account} from "./account";
import {AccountService} from "./account.service";
import {Router} from "@angular/router";

@Component({
  selector: 'create-account',
  template: `
    <div>
      <form action="/create_account" #createForm = "ngForm" (ngSubmit)="onSubmit()">
        <fieldset>
          <legend align="center">Account Information</legend>
          <label class="err_msg">{{error}}</label>
          <input type="hidden" name="create_account" value="true">
          <div>
            <label>First Name:</label>
            <input type="text" [(ngModel)]="acct.fname" name="fname" autofocus="on"  autocomplete="on" required>
          </div>
          <div>
            <label>Last Name:</label>
            <input type="text" [(ngModel)]="acct.lname" name="lname" autocomplete="on" required>
          </div>
          <div>
            <label>Email:</label>
            <input type="email" [(ngModel)]="acct.email" name="email" email autocomplete="on" required>
          </div>
          <div>
            <label>Password:</label>
            <input type="password" [(ngModel)]="acct.pwd" name="pwd" pattern=".{6,}" required>
          </div>
          <div>
            <label>Confirm Pwd:</label>
            <input type="password" [(ngModel)]="chkPwd"name="chkPwd" pattern=".{6,}" required>
          </div>
          <div>
            <button type="submit" [disabled]="!createForm.form.valid">Submit</button>
          </div>
          <a href="/login">Login</a>
        </fieldset>
      </form>
  </div>
  `,

  providers: [AccountService],
  styleUrls: ['./login.component.css']
})

export class CreateAccountComponent {
  @Input() acct: Account;
  @Input() chkPwd: string;
  error: string;

  constructor(
    private acctSvc: AccountService,
    private router: Router,
  ) {
    this.acct = new Account();
  }

  onSubmit(){
    if(this.passwordsMatch()) {
      this.acctSvc.createAccount(this.acct).subscribe(
        resp => this.processResponse(resp.status),
        err => this.processError(err.status)
      );
    }
  }

  passwordsMatch(): boolean {
    if(this.acct.pwd != this.chkPwd) {
      this.error = "Passwords Did Not Match";
      this.acct.pwd = "";
      this.chkPwd = ""
      return false;
    }
    return true;
  }

  processResponse(status: number){
    this.error = "Were Good!";
  }

  processError(status: number){
    switch (status) {
      case 400:
        this.error = "Malformed Request";
        break;
      case 401:
        this.error = "Email Address Already In Use";
        break;
      case 500:
        this.router.navigateByUrl('/app-error')
        break;
    }
    this.acct.clear()
    this.chkPwd = ""
  }
}
