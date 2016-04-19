import {Component} from 'angular2/core';
import {TopbarComponent} from '../shared/topbar/topbar.component'
import {SidebarComponent} from '../shared/sidebar/sidebar.component'


@Component({
  selector: 'my-app',
  templateUrl: 'app/components/app.component.html',
  directives: [
    TopbarComponent,
    SidebarComponent
  ]
})
export class AppComponent{
  title = 'My GoWA2 APP !!';
}
