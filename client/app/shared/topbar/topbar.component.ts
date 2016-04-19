import {Component, Input} from 'angular2/core'
import {LoginComponent} from '../login/login.component'

@Component({
  selector: 'my-topbar',
  templateUrl: 'app/shared/topbar/topbar.component.html',
  directives: [
    LoginComponent
  ]
})
export class TopbarComponent{
    @Input()
    title: string;
}
