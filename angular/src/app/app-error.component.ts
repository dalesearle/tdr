import { Component } from '@angular/core';

@Component({
  selector: 'app-error',
  template: `
    <div>
      <fieldset>
        <legend align="center">Application Error</legend>
        <label>The problem was logged and will be looked into.</label>
      </fieldset>
    </div>
  `,
  styles: [`
    * {
      font-family: Arial;
      font-size: 1em;
    }
    
    fieldset{
      border-radius: 12px;
      max-width: 300px;
      padding-bottom: 15px;
      padding-top: 15px;
    }
  `],
  providers: []
})

export class AppErrorComponent {
}
