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

//  срабатывает когда текстэреа меняет свое значение и теряет фокус: посылает в метод сам комментарий и айди
//  того, кому этот комментарий посвящен
$(".comment-area").change(function() {
    console.log("Event catched")
    var text = $(this).val()
    var liked_user_id = parseInt($(this).closest(".long_card").find(".user_id").text())
    console.log(JSON.stringify({ comment: text, liked_user_id: liked_user_id }, null, '\t'))
    $.ajax({
        url: '/comment',
        type: 'post',
        dataType: 'json',
        contentType: 'application/json',
        success: function (data) {
            console.log(data)
        },
        data: JSON.stringify({ comment: text, liked_user_id: liked_user_id }, null, '\t')
    });
})

//  жмешь на сердце с крестом - карточка пропадает запись в бд удаляется все счастливы LOVE-NO-MORE
$(".love-no-more").click(function() {
    var liked_user_id = parseInt($(this).closest(".long_card").find(".user_id").text())
    console.log("Going to remove user with id " + liked_user_id + " from favorites")
    $.ajax({
        url: '/dislike',
        type: 'post',
        dataType: 'json',
        contentType: 'application/json',
        success: function (data) {
            console.log(data)
            if (data == "OK") {
                $(`.user_id:contains('${liked_user_id}')`).closest(".long_card").remove()
            }
        },
        data: JSON.stringify(liked_user_id)
    });
})