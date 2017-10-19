import { Account } from './account';
import { Http, Response, Headers, RequestOptions } from '@angular/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Rx';


@Injectable()
export class AccountService {
  private headers = new Headers({'Content-Type': 'application/json'});

  constructor(private http: Http) {}

  login(acct: Account): Observable<Response> {
    let jsn = JSON.stringify(acct);
    let options = new RequestOptions({ headers: this.headers });
    return this.http.post("https://localhost:8081/login", jsn, options)
  }

  createAccount(acct: Account): Observable<Response> {
    let jsn = JSON.stringify(acct);
    let options = new RequestOptions({ headers: this.headers });
    return this.http.post("https://localhost:8081/create-account", jsn, options)
  }
}

