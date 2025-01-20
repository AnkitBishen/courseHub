app.controller('LoginController', function ($scope, $http) {
    
    $scope.cerdentials = {
        email: '',
        password: ''
    }
    $scope.login = function () {
        if ($scope.cerdentials.email != '' && $scope.cerdentials.password != '') {
            $http({
                method: 'POST',
                url:'http://localhost/GO/courseHub/public/requests/login.request.php',
                data: {cerdentials:$scope.cerdentials, authentication : true},
                headers: {
                    'Content-Type': 'application/json'
                }
            }).then(function (response) {
                if (response.data.status == 'success') {
                    // window.location.href = 'http://localhost/GO/courseHub/public/index.php';
                } else {
                    // alert(response.data.message);
                }
            }, function (response) {
                console.log(response);
            });
        }
    }
})