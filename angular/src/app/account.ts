export class Account {
  email: string;
  pwd: string;
  fname: string;
  lname: string;

  public clear() {
    this.email = "";
    this.pwd = "";
    this.fname = "";
    this.lname = "";
  }
}
