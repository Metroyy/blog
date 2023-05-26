$('#search').on('input', function () {
    var input = $(this).val()
    $.ajax({
        url: 'http://localhost:8080/Search',
        method: 'POST',
        data: { input: input },
        success: function (response) {
            if (input != "") {
                $('#search-results').css('display', 'block');
                var resultHTML = "";
                for (var i = 0; i < response.length; i++) {
                    var path = response[i].substring(5);
                    path = path.substring(0, path.length - 5);
                    resultHTML += "<a target='_blank' href='" + response[i] + "'>" + path + "</a><br>";
                }
                $('#search-results').html(resultHTML);
            }else {
                $('#search-results').css('display', 'none');
            }
        }
    })
})