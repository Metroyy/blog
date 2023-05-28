$('#search').on('input', function () {
    var input = $(this).val()
    $.ajax({
        url: 'http://localhost:4005/Search',
        method: 'POST',
        data: { input: input },
        success: function (response) {
            if (input != "") {
                var resultHTML = "";
                for (var i = 0; i < response.length; i++) {
                    var path = response[i].substring(5);
                    path = path.substring(0, path.length - 5);
                    resultHTML += "<a target='_blank' href='" + response[i] + "'>" + path + "</a><br>";
                }
                $('#search-results').html(resultHTML);
            }
        }
    })
})