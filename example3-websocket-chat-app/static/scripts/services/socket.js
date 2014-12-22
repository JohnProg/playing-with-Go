'use strict';

angular.module('chatWebApp')
    .factory('socket', ['socketFactory', function (socketFactory) {
        var socket = socketFactory();
        socket.forward('error');
        return socket;
    }]);