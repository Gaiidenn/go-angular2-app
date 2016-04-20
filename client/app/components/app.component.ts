import {Component} from 'angular2/core';
import {RouteConfig, ROUTER_DIRECTIVES, ROUTER_PROVIDERS} from 'angular2/router';
import {TopbarComponent} from '../shared/topbar/topbar.component';
import {SidebarComponent} from '../shared/sidebar/sidebar.component';
import {DashboardComponent} from '../components/dashboard/dashboard.component';
import {TodoComponent} from '../components/todo/todo.component';

@Component({
    selector: 'my-app',
    templateUrl: 'app/components/app.component.html',
    directives: [
        ROUTER_DIRECTIVES,
        TopbarComponent,
        SidebarComponent
    ],
    providers: [
        ROUTER_PROVIDERS
    ]
})
@RouteConfig([
    {
        path: '/dashboard',
        name: 'Dashboard',
        component: DashboardComponent,
        useAsDefault: true
    },
    {
        path: '/todo',
        name: 'Todo',
        component: TodoComponent
    }
])
export class AppComponent{
    title = 'My GoWA2 APP !!';
}
