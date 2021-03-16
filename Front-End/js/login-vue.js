var app = new Vue({
    el : "#container",
    data : {
        isRegister:false,
        // 图片路径请自行更改
        imgSrc : ["img/1.jpg","img/2.jpg","img/3.jpg","img/4.jpg",
    ],
        imgIndex : 0,
    },
    created: function () {
        setInterval(this.lantenSlide, 2000);
    },

    methods : {
        selectLogin:function(){
            // alert("login")
            this.isRegister = false;
        },
        selectRegister:function(){
            // alert("register")
            this.isRegister = true;
        },
        lantenSlide : function(){
            this.imgIndex = ((this.imgIndex+1) % this.imgSrc.length)
            console.log(this.imgIndex);
        },

    }
});   