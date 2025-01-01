
$(document).ready(function () {

    $(".sendAjax").on("click", function () {
        selectValue = $("#citySelect").val()

        if (selectValue) {
            $.ajax({
                url: "/api/process",
                type: "POST",
                contentType: "application/json",
                data: JSON.stringify({City: selectValue}),
                dataType: "json",
                success: function (response) {
                    $("#countryResponce").html("Country: " + response.location.country)
                    $("#cityResponce").html("City: " + response.location.name)
                    $("#tempResponce").html("Temperature: " + response.current.temp_c)
                    $(".responceDiv").show()
                },
                error: function (error) {
                    console.error("Error:", error);
                }
            });
        }
    });
});


