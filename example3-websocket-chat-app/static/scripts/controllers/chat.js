'use strict';

angular.module('chatWebApp')
    .controller('ChatCtrl', ['$scope', 'socket', function ($scope, socket) {
        $scope.messages = [];
        $scope.newMessage = '';
        $scope.username = false;
        $scope.inputUsername = '';
        $scope.glued = true;
        var Notification = window.Notification || window.mozNotification || window.webkitNotification;

        socket.forward('message', $scope);
        $scope.$on('socket:message', function (ev, data) {
            if ($scope.messages.length > 100) {
                $scope.messages.splice(0, 1);
            }
            var msg = JSON.parse(data);
            $scope.messages.push(msg);

            var hidden = false;
            if (typeof document.hidden !== "undefined") {
                hidden = "hidden";
            } else if (typeof document.mozHidden !== "undefined") {
                hidden = "mozHidden";
            } else if (typeof document.msHidden !== "undefined") {
                hidden = "msHidden";
            } else if (typeof document.webkitHidden !== "undefined") {
                hidden = "webkitHidden";
            }

            // $scope.username is not set if the user didn't provide a name and thus didn't display the chat window
            // document[hidden] is true if the page is minimized or tabbed-out — details vary by browser
            if ($scope.username && document[hidden] && msg.type == 'message') {
                var instance = new Notification(
                    msg.username + " says:", {
                         body: msg.message
                     }
                );
            }
        });

        $scope.sendMessage = function () {
            socket.emit('send_message', $scope.newMessage);
            $scope.messages.push($scope.newMessage);
            $scope.newMessage = '';
        };

        $scope.setUsername = function () {
            $scope.username = $scope.inputUsername;
            socket.emit('joined_message', $scope.username);
            // setUsername is called once and can be regarded as "login"
            Notification.requestPermission(function (permission) {
            });
        };
    }]);
