import { Component } from '@angular/core';

@Component({
  selector: 'menu',
  template: `
    <nav>
      <a href="/raids">Raids</a>
      <a href="/infest">Infest</a>
    </nav>`,
  styles: [`
    * {
      color: black;
      font-family: Arial;
      font-size: 1em;
    }
    
    a{
      display: inline-block
      ;
      margin: 3px;
      padding: 3px;
      test-decoration: none;
      border: 1px solid black;
      border-radius:6px;
      text-align: center;
      width: 15%;
    }
    
    a:link {
      text-decoration: none;
    }
    
    /* visited link */
    a:visited {
        text-decoration: none;
    }
    
    /* mouse over link */
    a:hover {
        text-decoration: none;
        font-weight: bold;
    }
    
    /* selected link */
    a:active {
        font-weight: bold;
    }
  `],
  providers: []
})

export class MenuComponent{

}
