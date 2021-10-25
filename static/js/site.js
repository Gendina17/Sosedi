//  клик по иконке отправляет пост запрос с айди кликнутого юзера на сервер и скрывает саму иконку
$(".icon_wraper").click(function() {
    $(this).attr("hidden", true)
    var liked_user_id = parseInt($(this).parent().find(".user_id").text())
    $.ajax({
        url: '/like',
        type: 'post',
        dataType: 'json',
        contentType: 'application/json',
        success: function (data) {
            console.log(data)
        },
        data: JSON.stringify(liked_user_id)
    });
})

//  при загрузке страницы отправляется пост запрос на сервер чтобы получить информацию о текущей сессии
//  в зависимости от ответа либо не делается ничего, либо делаются видимыми иконки лайка еще не лайкнутых юзеров
//  короче я пошел по легкому пути сделал через js это все, поэтому если захочешь запариться можешь попробовать
//  через шаблоны это сделать
$( document ).ready(function() {
    if ($(location).attr("href") == "http://localhost:8080/")
    {
        $.ajax({
            url: '/check_session',
            type: 'post',
            dataType: 'json',
            contentType: 'application/json',
            success: function (data) {
                console.log(data)
                if (data != "Inactive") {
                    $("main .icon_wraper").each(function() {
                        var card_user_id = parseInt($(this).parent().find(".user_id").text())
                        if (!data.includes(card_user_id)) {
                            $(this).attr("hidden", false)
                        }
                    })
                }
            }
        })
    }   
})