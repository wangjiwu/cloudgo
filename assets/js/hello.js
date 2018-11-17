$(document).ready(function() {
    $.ajax({
        url: "/api/test"
    }).then(function(data) {
       $('.greeting').append(data.WELCOME);
    });
});
