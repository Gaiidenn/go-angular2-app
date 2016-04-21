import {Component} from 'angular2/core';
import {RouteConfig, ROUTER_DIRECTIVES, ROUTER_PROVIDERS, Router} from 'angular2/router';
import {$WebSocket} from 'angular2-websocket/angular2-websocket';
import {DashboardComponent} from '../components/dashboard/dashboard.component';
import {TodoComponent} from '../components/todo/todo.component';

@Component({
    selector: 'my-app',
    templateUrl: 'app/components/app.component.html',
    directives: [
        ROUTER_DIRECTIVES
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
    private _ws: $WebSocket;

    constructor(
        private _router: Router
    ) {
        this._ws = new $WebSocket("ws://localhost:8080/");
        let cb = function(message: any) {
            alert(message.data);
        }
        this._ws.onMessage(cb, null);
    }

    sendMessage(message: string) {
        this._ws.send(message);
        //alert('message "' + message + '" sent!');
    }
}
