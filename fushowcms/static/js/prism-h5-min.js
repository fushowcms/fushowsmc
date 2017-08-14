/*
 * Copyright (c) 2016—2017 www.fushow.cn, All rights reserved.
 */
!
function a(b, c, d) {
    function e(g, h) {
        if (!c[g]) {
            if (!b[g]) {
                var i = "function" == typeof require && require;
                if (!h && i) return i(g, !0);
                if (f) return f(g, !0);
                var j = new Error("Cannot find module '" + g + "'");
                throw j.code = "MODULE_NOT_FOUND",
                j
            }
            var k = c[g] = {
                exports: {}
            };
            b[g][0].call(k.exports,
            function(a) {
                var c = b[g][1][a];
                return e(c ? c: a)
            },
            k, k.exports, a, b, c, d)
        }
        return c[g].exports
    }
    for (var f = "function" == typeof require && require,
    g = 0; g < d.length; g++) e(d[g]);
    return e
} ({
    1 : [function(a, b, c) {
        b.exports = {
            domain: "g.alicdn.com",
            flashVersion: "1.2.12",
            h5Version: "1.5.7",
            logReportTo: "//videocloud.cn-hangzhou.log.aliyuncs.com/logstores/player/track"
        }
    },
    {}],
    2 : [function(a, b, c) {
        var d = a("./player/player"),
        e = a("./lib/dom"),
        f = a("./lib/ua"),
        g = a("./lib/object"),
        h = a("./config"),
        i = function(a) {
            var b, c = a.id;
            if ("string" == typeof c) {
                if (0 === c.indexOf("#") && (c = c.slice(1)), i.players[c]) return i.players[c];
                b = e.el(c)
            } else b = c;
            if (!b || !b.nodeName) throw new TypeError("没有为播放器指定容器");
            var h = g.merge(i.defaultOpt, a);
            if (a.isLive && (h.skinLayout = [{
                name: "bigPlayButton",
                align: "blabs",
                x: 30,
                y: 80
            },
            {
                name: "controlBar",
                align: "blabs",
                x: 0,
                y: 0,
                children: [{
                    name: "liveDisplay",
                    align: "tlabs",
                    x: 15,
                    y: 25
                },
                {
                    name: "fullScreenButton",
                    align: "tr",
                    x: 20,
                    y: 25
                },
                {
                    name: "volume",
                    align: "tr",
                    x: 20,
                    y: 25
                }]
            }]), f.IS_IOS) for (var j = 0; j < h.skinLayout.length; j++) if ("controlBar" == h.skinLayout[j].name) for (var k = h.skinLayout[j], l = 0; l < k.children.length; l++) if ("volume" == k.children[l].name) {
                k.children.splice(l, 1);
                break
            }
            if (h.width && (b.style.width = h.width), h.height) {
                var m = h.height.indexOf("%");
                if (m > 0) {
                    var n = window.screen.height,
                    o = h.height.replace("%", "");
                    if (isNaN(o)) b.style.height = h.height;
                    else {
                        var p = 9 * n * parseInt(o) / 1e3;
                        b.style.height = String(p % 2 ? p + 1 : p) + "px"
                    }
                } else b.style.height = h.height
            }
            return b.player || new d(b, h)
        },
        j = window.prismplayer = i;
        i.players = {},
        i.defaultOpt = {
            preload: !1,
            autoplay: !0,
            useNativeControls: !1,
            width: "100%",
            height: "300px",
            cover: "",
            from: "prism_aliyun",
            trackLog: !0,
            isLive: !1,
            playsinline: !1,
            showBarTime: 5e3,
            rePlay: !1,
            skinRes: "//" + h.domain + "/de/prismplayer-flash/" + h.flashVersion + "/atlas/defaultSkin",
            skinLayout: [{
                name: "bigPlayButton",
                align: "blabs",
                x: 30,
                y: 80
            },
            {
                name: "controlBar",
                align: "blabs",
                x: 0,
                y: 0,
                children: [{
                    name: "progress",
                    align: "tlabs",
                    x: 0,
                    y: 0
                },
                {
                    name: "playButton",
                    align: "tl",
                    x: 15,
                    y: 26
                },
                {
                    name: "timeDisplay",
                    align: "tl",
                    x: 10,
                    y: 24
                },
                {
                    name: "fullScreenButton",
                    align: "tr",
                    x: 20,
                    y: 25
                },
                {
                    name: "volume",
                    align: "tr",
                    x: 20,
                    y: 25
                }]
            }]
        },
        "function" == typeof define && define.amd ? define([],
        function() {
            return j
        }) : "object" == typeof c && "object" == typeof b && (b.exports = j)
    },
    {
        "./config": 1,
        "./lib/dom": 5,
        "./lib/object": 10,
        "./lib/ua": 12,
        "./player/player": 16
    }],
    3 : [function(a, b, c) {
        b.exports.get = function(a) {
            for (var b = a + "",
            c = document.cookie.split(";"), d = 0; d < c.length; d++) {
                var e = c[d].trim();
                if (0 == e.indexOf(b)) return unescape(e.substring(b.length + 1, e.length))
            }
            return ""
        },
        b.exports.set = function(a, b, c) {
            var d = new Date;
            d.setTime(d.getTime() + 24 * c * 60 * 60 * 1e3);
            var e = "expires=" + d.toGMTString();
            document.cookie = a + "=" + escape(b) + "; " + e
        }
    },
    {}],
    4 : [function(a, b, c) {
        var d = a("./object");
        b.exports.cache = {},
        b.exports.guid = function(a, b) {
            var c, d = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz".split(""),
            e = [];
            if (b = b || d.length, a) for (c = 0; c < a; c++) e[c] = d[0 | Math.random() * b];
            else {
                var f;
                for (e[8] = e[13] = e[18] = e[23] = "-", e[14] = "4", c = 0; c < 36; c++) e[c] || (f = 0 | 16 * Math.random(), e[c] = d[19 == c ? 3 & f | 8 : f])
            }
            return e.join("")
        },
        b.exports.expando = "vdata" + (new Date).getTime(),
        b.exports.getData = function(a) {
            var c = a[b.exports.expando];
            return c || (c = a[b.exports.expando] = b.exports.guid(), b.exports.cache[c] = {}),
            b.exports.cache[c]
        },
        b.exports.hasData = function(a) {
            var c = a[b.exports.expando];
            return ! (!c || d.isEmpty(b.exports.cache[c]))
        },
        b.exports.removeData = function(a) {
            var c = a[b.exports.expando];
            if (c) {
                delete b.exports.cache[c];
                try {
                    delete a[b.exports.expando]
                } catch(c) {
                    a.removeAttribute ? a.removeAttribute(b.exports.expando) : a[b.exports.expando] = null
                }
            }
        }
    },
    {
        "./object": 10
    }],
    5 : [function(a, b, c) {
        var d = a("./object");
        b.exports.el = function(a) {
            return document.getElementById(a)
        },
        b.exports.createEl = function(a, b) {
            var c;
            return a = a || "div",
            b = b || {},
            c = document.createElement(a),
            d.each(b,
            function(a, b) {
                a.indexOf("aria-") !== -1 || "role" == a ? c.setAttribute(a, b) : c[a] = b
            }),
            c
        },
        b.exports.addClass = function(a, b) { (" " + a.className + " ").indexOf(" " + b + " ") == -1 && (a.className = "" === a.className ? b: a.className + " " + b)
        },
        b.exports.removeClass = function(a, b) {
            var c, d;
            if (a.className.indexOf(b) != -1) {
                for (c = a.className.split(" "), d = c.length - 1; d >= 0; d--) c[d] === b && c.splice(d, 1);
                a.className = c.join(" ")
            }
        },
        b.exports.getElementAttributes = function(a) {
            var b, c, d, e, f;
            if (b = {},
            c = ",autoplay,controls,loop,muted,default,", a && a.attributes && a.attributes.length > 0) {
                d = a.attributes;
                for (var g = d.length - 1; g >= 0; g--) e = d[g].name,
                f = d[g].value,
                "boolean" != typeof a[e] && c.indexOf("," + e + ",") === -1 || (f = null !== f),
                b[e] = f
            }
            return b
        },
        b.exports.insertFirst = function(a, b) {
            b.firstChild ? b.insertBefore(a, b.firstChild) : b.appendChild(a)
        },
        b.exports.blockTextSelection = function() {
            document.body.focus(),
            document.onselectstart = function() {
                return ! 1
            }
        },
        b.exports.unblockTextSelection = function() {
            document.onselectstart = function() {
                return ! 0
            }
        },
        b.exports.css = function(a, b, c) {
            return !! a.style && (b && c ? (a.style[b] = c, !0) : c || "string" != typeof b ? !c && "object" == typeof b && (d.each(b,
            function(b, c) {
                a.style[b] = c
            }), !0) : a.style[b])
        }
    },
    {
        "./object": 10
    }],
    6 : [function(a, b, c) {
        function d(a, b, c, d) {
            e.each(c,
            function(c) {
                a(b, c, d)
            })
        }
        var e = a("./object"),
        f = a("./data");
        b.exports.on = function(a, c, g) {
            if (e.isArray(c)) return d(b.exports.on, a, c, g);
            var h = f.getData(a);
            h.handlers || (h.handlers = {}),
            h.handlers[c] || (h.handlers[c] = []),
            g.guid || (g.guid = f.guid()),
            h.handlers[c].push(g),
            h.dispatcher || (h.disabled = !1, h.dispatcher = function(c) {
                if (!h.disabled) {
                    c = b.exports.fixEvent(c);
                    var d = h.handlers[c.type];
                    if (d) for (var e = d.slice(0), f = 0, g = e.length; f < g && !c.isImmediatePropagationStopped(); f++) e[f].call(a, c)
                }
            }),
            1 == h.handlers[c].length && (a.addEventListener ? a.addEventListener(c, h.dispatcher, !1) : a.attachEvent && a.attachEvent("on" + c, h.dispatcher))
        },
        b.exports.off = function(a, c, g) {
            if (f.hasData(a)) {
                var h = f.getData(a);
                if (h.handlers) {
                    if (e.isArray(c)) return d(b.exports.off, a, c, g);
                    var i = function(c) {
                        h.handlers[c] = [],
                        b.exports.cleanUpEvents(a, c)
                    };
                    if (c) {
                        var j = h.handlers[c];
                        if (j) {
                            if (!g) return void i(c);
                            if (g.guid) for (var k = 0; k < j.length; k++) j[k].guid === g.guid && j.splice(k--, 1);
                            b.exports.cleanUpEvents(a, c)
                        }
                    } else for (var l in h.handlers) i(l)
                }
            }
        },
        b.exports.cleanUpEvents = function(a, b) {
            var c = f.getData(a);
            0 === c.handlers[b].length && (delete c.handlers[b], a.removeEventListener ? a.removeEventListener(b, c.dispatcher, !1) : a.detachEvent && a.detachEvent("on" + b, c.dispatcher)),
            e.isEmpty(c.handlers) && (delete c.handlers, delete c.dispatcher, delete c.disabled),
            e.isEmpty(c) && f.removeData(a)
        },
        b.exports.fixEvent = function(a) {
            function b() {
                return ! 0
            }
            function c() {
                return ! 1
            }
            if (!a || !a.isPropagationStopped) {
                var d = a || window.event;
                a = {};
                for (var e in d)"layerX" !== e && "layerY" !== e && "keyboardEvent.keyLocation" !== e && ("returnValue" == e && d.preventDefault || (a[e] = d[e]));
                if (a.target || (a.target = a.srcElement || document), a.relatedTarget = a.fromElement === a.target ? a.toElement: a.fromElement, a.preventDefault = function() {
                    d.preventDefault && d.preventDefault(),
                    a.returnValue = !1,
                    a.isDefaultPrevented = b,
                    a.defaultPrevented = !0
                },
                a.isDefaultPrevented = c, a.defaultPrevented = !1, a.stopPropagation = function() {
                    d.stopPropagation && d.stopPropagation(),
                    a.cancelBubble = !0,
                    a.isPropagationStopped = b
                },
                a.isPropagationStopped = c, a.stopImmediatePropagation = function() {
                    d.stopImmediatePropagation && d.stopImmediatePropagation(),
                    a.isImmediatePropagationStopped = b,
                    a.stopPropagation()
                },
                a.isImmediatePropagationStopped = c, null != a.clientX) {
                    var f = document.documentElement,
                    g = document.body;
                    a.pageX = a.clientX + (f && f.scrollLeft || g && g.scrollLeft || 0) - (f && f.clientLeft || g && g.clientLeft || 0),
                    a.pageY = a.clientY + (f && f.scrollTop || g && g.scrollTop || 0) - (f && f.clientTop || g && g.clientTop || 0)
                }
                a.which = a.charCode || a.keyCode,
                null != a.button && (a.button = 1 & a.button ? 0 : 4 & a.button ? 1 : 2 & a.button ? 2 : 0)
            }
            return a
        },
        b.exports.trigger = function(a, c) {
            var d = f.hasData(a) ? f.getData(a) : {},
            e = a.parentNode || a.ownerDocument;
            if ("string" == typeof c) {
                var g = null;
                a.paramData && (g = a.paramData, a.paramData = null, a.removeAttribute(g)),
                c = {
                    type: c,
                    target: a,
                    paramData: g
                }
            }
            if (c = b.exports.fixEvent(c), d.dispatcher && d.dispatcher.call(a, c), e && !c.isPropagationStopped() && c.bubbles !== !1) b.exports.trigger(e, c);
            else if (!e && !c.defaultPrevented) {
                var h = f.getData(c.target);
                c.target[c.type] && (h.disabled = !0, "function" == typeof c.target[c.type] && c.target[c.type](), h.disabled = !1)
            }
            return ! c.defaultPrevented
        },
        b.exports.one = function(a, c, g) {
            if (e.isArray(c)) return d(b.exports.one, a, c, g);
            var h = function() {
                b.exports.off(a, c, h),
                g.apply(this, arguments)
            };
            h.guid = g.guid = g.guid || f.guid(),
            b.exports.on(a, c, h)
        }
    },
    {
        "./data": 4,
        "./object": 10
    }],
    7 : [function(a, b, c) {
        var d = a("./data");
        b.exports.bind = function(a, b, c) {
            b.guid || (b.guid = d.guid());
            var e = function() {
                return b.apply(a, arguments)
            };
            return e.guid = c ? c + "_" + b.guid: b.guid,
            e
        }
    },
    {
        "./data": 4
    }],
    8 : [function(a, b, c) {
        var d = a("./url");
        b.exports.get = function(a, b, c, e) {
            var f, g, h, i, j;
            c = c ||
            function() {},
            "undefined" == typeof XMLHttpRequest && (window.XMLHttpRequest = function() {
                try {
                    return new window.ActiveXObject("Msxml2.XMLHTTP.6.0")
                } catch(a) {}
                try {
                    return new window.ActiveXObject("Msxml2.XMLHTTP.3.0")
                } catch(a) {}
                try {
                    return new window.ActiveXObject("Msxml2.XMLHTTP")
                } catch(a) {}
                throw new Error("This browser does not support XMLHttpRequest.")
            }),
            g = new XMLHttpRequest,
            h = d.parseUrl(a),
            i = window.location,
            j = h.protocol + h.host !== i.protocol + i.host,
            !j || !window.XDomainRequest || "withCredentials" in g ? (f = "file:" == h.protocol || "file:" == i.protocol, g.onreadystatechange = function() {
                4 === g.readyState && (200 === g.status || f && 0 === g.status ? b(g.responseText) : c(g.responseText))
            }) : (g = new window.XDomainRequest, g.onload = function() {
                b(g.responseText)
            },
            g.onerror = c, g.onprogress = function() {},
            g.ontimeout = c);
            try {
                g.open("GET", a, !0),
                e && (g.withCredentials = !0)
            } catch(a) {
                return void c(a)
            }
            try {
                g.send()
            } catch(a) {
                c(a)
            }
        },
        b.exports.jsonp = function(a, b, c) {
            var d = "jsonp_callback_" + Math.round(1e5 * Math.random()),
            e = document.createElement("script");
            e.src = a + (a.indexOf("?") >= 0 ? "&": "?") + "callback=" + d,
            e.onerror = function() {
                delete window[d],
                document.body.removeChild(e),
                c()
            },
            e.onload = function() {
                setTimeout(function() {
                    window[d] && (delete window[d], document.body.removeChild(e))
                },
                0)
            },
            window[d] = function(a) {
                delete window[d],
                document.body.removeChild(e),
                b(a)
            },
            document.body.appendChild(e)
        }
    },
    {
        "./url": 13
    }],
    9 : [function(a, b, c) {
        var d = a("./dom");
        b.exports.render = function(a, b) {
            var c = b.align ? b.align: "tl",
            e = b.x ? b.x: 0,
            f = b.y ? b.y: 0;
            "tl" === c ? d.css(a, {
                float: "left",
                "margin-left": e + "px",
                "margin-top": f + "px"
            }) : "tr" === c ? d.css(a, {
                float: "right",
                "margin-right": e + "px",
                "margin-top": f + "px"
            }) : "tlabs" === c ? d.css(a, {
                position: "absolute",
                left: e + "px",
                top: f + "px"
            }) : "trabs" === c ? d.css(a, {
                position: "absolute",
                right: e + "px",
                top: f + "px"
            }) : "blabs" === c ? d.css(a, {
                position: "absolute",
                left: e + "px",
                bottom: f + "px"
            }) : "brabs" === c ? d.css(a, {
                position: "absolute",
                right: e + "px",
                bottom: f + "px"
            }) : "cc" === c && d.css(a, {
                position: "absolute",
                left: "50%",
                top: "50%",
                "margin-top": a.offsetHeight / -2 + "px",
                "margin-left": a.offsetWidth / -2 + "px"
            })
        }
    },
    {
        "./dom": 5
    }],
    10 : [function(a, b, c) {
        var d = Object.prototype.hasOwnProperty;
        b.exports.create = Object.create ||
        function(a) {
            function b() {}
            return b.prototype = a,
            new b
        },
        b.exports.isArray = function(a) {
            return "[object Array]" === Object.prototype.toString.call(arg)
        },
        b.exports.isEmpty = function(a) {
            for (var b in a) if (null !== a[b]) return ! 1;
            return ! 0
        },
        b.exports.each = function(a, c, e) {
            if (b.exports.isArray(a)) for (var f = 0,
            g = a.length; f < g && c.call(e || this, a[f], f) !== !1; ++f);
            else for (var h in a) if (d.call(a, h) && c.call(e || this, h, a[h]) === !1) break;
            return a
        },
        b.exports.merge = function(a, b) {
            if (!b) return a;
            for (var c in b) d.call(b, c) && (a[c] = b[c]);
            return a
        },
        b.exports.deepMerge = function(a, c) {
            var e, f, g;
            a = b.exports.copy(a);
            for (e in c) d.call(c, e) && (f = a[e], g = c[e], b.exports.isPlain(f) && b.exports.isPlain(g) ? a[e] = b.exports.deepMerge(f, g) : a[e] = c[e]);
            return a
        },
        b.exports.copy = function(a) {
            return b.exports.merge({},
            a)
        },
        b.exports.isPlain = function(a) {
            return !! a && "object" == typeof a && "[object Object]" === a.toString() && a.constructor === Object
        },
        b.exports.isArray = Array.isArray ||
        function(a) {
            return "[object Array]" === Object.prototype.toString.call(a)
        },
        b.exports.unescape = function(a) {
            return a.replace(/&([^;]+);/g,
            function(a, b) {
                return {
                    amp: "&",
                    lt: "<",
                    gt: ">",
                    quot: '"',
                    "#x27": "'",
                    "#x60": "`"
                } [b.toLowerCase()] || a
            })
        }
    },
    {}],
    11 : [function(a, b, c) {
        var d = a("./object"),
        e = function() {},
        e = function() {};
        e.extend = function(a) {
            var b, c;
            a = a || {},
            b = a.init || a.init || this.prototype.init || this.prototype.init ||
            function() {},
            c = function() {
                b.apply(this, arguments)
            },
            c.prototype = d.create(this.prototype),
            c.prototype.constructor = c,
            c.extend = e.extend,
            c.create = e.create;
            for (var f in a) a.hasOwnProperty(f) && (c.prototype[f] = a[f]);
            return c
        },
        e.create = function() {
            var a = d.create(this.prototype);
            return this.apply(a, arguments),
            a
        },
        b.exports = e
    },
    {
        "./object": 10
    }],
    12 : [function(a, b, c) {
        if (b.exports.USER_AGENT = navigator.userAgent, b.exports.IS_IPHONE = /iPhone/i.test(b.exports.USER_AGENT), b.exports.IS_IPAD = /iPad/i.test(b.exports.USER_AGENT), b.exports.IS_IPOD = /iPod/i.test(b.exports.USER_AGENT), b.exports.IS_MAC = /mac/i.test(b.exports.USER_AGENT), b.exports.IS_SAFARI = /Safari/i.test(b.exports.USER_AGENT), b.exports.IS_CHROME = /Chrome/i.test(b.exports.USER_AGENT), b.exports.IS_FIREFOX = /Firefox/i.test(b.exports.USER_AGENT), document.all) {
            var d = new ActiveXObject("ShockwaveFlash.ShockwaveFlash");
            d ? b.exports.HAS_FLASH = !0 : b.exports.HAS_FLASH = !1
        } else if (navigator.plugins && navigator.plugins.length > 0) {
            var d = navigator.plugins["Shockwave Flash"];
            d ? b.exports.HAS_FLASH = !0 : b.exports.HAS_FLASH = !1
        } else b.exports.HAS_FLASH = !1;
        b.exports.IS_MAC_SAFARI = b.exports.IS_MAC && b.exports.IS_SAFARI && !b.exports.IS_CHROME && !b.exports.HAS_FLASH,
        b.exports.IS_IOS = b.exports.IS_IPHONE || b.exports.IS_IPAD || b.exports.IS_IPOD || b.exports.IS_MAC_SAFARI,
        b.exports.IOS_VERSION = function() {
            var a = b.exports.USER_AGENT.match(/OS (\d+)_/i);
            if (a && a[1]) return a[1]
        } (),
        b.exports.IS_ANDROID = /Android/i.test(b.exports.USER_AGENT),
        b.exports.ANDROID_VERSION = function() {
            var a, c, d = b.exports.USER_AGENT.match(/Android (\d+)(?:\.(\d+))?(?:\.(\d+))*/i);
            return d ? (a = d[1] && parseFloat(d[1]), c = d[2] && parseFloat(d[2]), a && c ? parseFloat(d[1] + "." + d[2]) : a ? a: null) : null
        } (),
        b.exports.IS_OLD_ANDROID = b.exports.IS_ANDROID && /webkit/i.test(b.exports.USER_AGENT) && b.exports.ANDROID_VERSION < 2.3,
        b.exports.TOUCH_ENABLED = !!("ontouchstart" in window || window.DocumentTouch && document instanceof window.DocumentTouch),
        b.exports.IS_MOBILE = b.exports.IS_IOS || b.exports.IS_ANDROID,
        b.exports.IS_H5 = b.exports.IS_MOBILE || !b.exports.HAS_FLASH,
        b.exports.IS_PC = !b.exports.IS_H5
    },
    {}],
    13 : [function(a, b, c) {
        var d = a("./dom");
        b.exports.getAbsoluteURL = function(a) {
            return a.match(/^https?:\/\//) || (a = d.createEl("div", {
                innerHTML: '<a href="' + a + '">x</a>'
            }).firstChild.href),
            a
        },
        b.exports.parseUrl = function(a) {
            var b, c, e, f, g;
            f = ["protocol", "hostname", "port", "pathname", "search", "hash", "host"],
            c = d.createEl("a", {
                href: a
            }),
            e = "" === c.host && "file:" !== c.protocol,
            e && (b = d.createEl("div"), b.innerHTML = '<a href="' + a + '"></a>', c = b.firstChild, b.setAttribute("style", "display:none; position:absolute;"), document.body.appendChild(b)),
            g = {};
            for (var h = 0; h < f.length; h++) g[f[h]] = c[f[h]];
            return e && document.body.removeChild(b),
            g
        }
    },
    {
        "./dom": 5
    }],
    14 : [function(a, b, c) {
        b.exports.formatTime = function(a) {
            var b, c, d, e = Math.round(a);
            return b = Math.floor(e / 3600),
            e %= 3600,
            c = Math.floor(e / 60),
            d = e % 60,
            !(b === 1 / 0 || isNaN(b) || c === 1 / 0 || isNaN(c) || d === 1 / 0 || isNaN(d)) && (b = b >= 10 ? b: "0" + b, c = c >= 10 ? c: "0" + c, d = d >= 10 ? d: "0" + d, ("00" === b ? "": b + ":") + c + ":" + d)
        },
        b.exports.parseTime = function(a) {
            var b = a.split(":"),
            c = 0,
            d = 0,
            e = 0;
            return 3 === b.length ? (c = b[0], d = b[1], e = b[2]) : 2 === b.length ? (d = b[0], e = b[1]) : 1 === b.length && (e = b[0]),
            c = parseInt(c, 10),
            d = parseInt(d, 10),
            e = Math.ceil(parseFloat(e)),
            3600 * c + 60 * d + e
        }
    },
    {}],
    15 : [function(a, b, c) {
        var d, e = a("../lib/oo"),
        f = a("../lib/object"),
        g = a("../lib/cookie"),
        h = a("../lib/data"),
        i = a("../lib/io"),
        j = a("../lib/ua"),
        k = a("../config"),
        l = {
            INIT: 1001,
            CLOSE: 1002,
            PLAY: 2001,
            STOP: 2002,
            PAUSE: 2003,
            RECOVER: 2010,
            SEEK: 2004,
            SEEK_END: 2011,
            FULLSREEM: 2005,
            QUITFULLSCREEM: 2006,
            UNDERLOAD: 3002,
            LOADED: 3001,
            RESOLUTION: 2007,
            RESOLUTION_DONE: 2009,
            HEARTBEAT: 9001,
            ERROR: 4001
        },
        m = e.extend({
            init: function(a, b) {
                this.player = a;
                var c = this.player.getOptions(),
                d = "1",
                e = b.from ? b.from: "prism_aliyun",
                f = c.isLive ? "prism_live": "prism_vod",
                g = "pc";
                j.IS_IPAD ? g = "pad": j.IS_IPHONE ? g = "iphone": j.IS_ANDROID && (g = "andorid");
                var h = j.IS_PC ? "pc_h5": "h5",
                i = k.h5Version,
                l = this._getUuid(),
                m = c.source ? encodeURIComponent(c.source) : b.video_id,
                n = "0",
                o = this.sessionId,
                p = "0",
                q = "0",
                r = "custom",
                s = "0.0.0.0",
                t = (new Date).getTime();
                this.opt = {
                    APIVersion: "0.6.0",
                    lv: d,
                    b: e,
                    lm: f,
                    t: g,
                    m: h,
                    pv: i,
                    uuid: l,
                    v: m,
                    u: n,
                    s: o,
                    e: p,
                    args: q,
                    d: r,
                    cdn_ip: s,
                    ct: t
                },
                this.bindEvent()
            },
            updateVideoInfo: function(a) {
                var b = this.player.getOptions(),
                c = "1",
                d = a.from ? a.from: "prism_aliyun",
                e = b.isLive ? "prism_live": "prism_vod",
                f = "pc";
                j.IS_IPAD ? f = "pad": j.IS_IPHONE ? f = "iphone": j.IS_ANDROID && (f = "andorid");
                var g = j.IS_PC ? "pc_h5": "h5",
                h = k.h5Version,
                i = this._getUuid(),
                l = b.source ? encodeURIComponent(b.source) : a.video_id,
                m = "0",
                n = this.sessionId,
                o = "0",
                p = "0",
                q = "custom",
                r = "0.0.0.0",
                s = (new Date).getTime();
                this.opt = {
                    APIVersion: "0.6.0",
                    lv: c,
                    b: d,
                    lm: e,
                    t: f,
                    m: g,
                    pv: h,
                    uuid: i,
                    v: l,
                    u: m,
                    s: n,
                    e: o,
                    args: p,
                    d: q,
                    cdn_ip: r,
                    ct: s
                }
            },
            bindEvent: function() {
                var a = this;
                this.player.on("init",
                function() {
                    a._onPlayerInit()
                }),
                window.addEventListener("beforeunload",
                function() {
                    a._onPlayerClose()
                }),
                this.player.on("ready",
                function() {
                    a._onPlayerReady()
                }),
                this.player.on("ended",
                function() {
                    a._onPlayerFinish()
                }),
                this.player.on("play",
                function() {
                    a._onPlayerPlay()
                }),
                this.player.on("pause",
                function() {
                    a._onPlayerPause()
                }),
                this.player.on("seekStart",
                function(b) {
                    a._onPlayerSeekStart(b)
                }),
                this.player.on("seekEnd",
                function(b) {
                    a._onPlayerSeekEnd(b)
                }),
                this.player.on("waiting",
                function() {
                    a._onPlayerLoaded()
                }),
                this.player.on("canplaythrough",
                function() {
                    a._onPlayerUnderload()
                }),
                this.player.on("error",
                function() {
                    a._onPlayerError()
                }),
                d = setInterval(function() {
                    2 === a.player.readyState() || 3 === a.player.readyState() ? a._onPlayerLoaded() : 4 === a.player.readyState() && a._onPlayerUnderload()
                },
                100)
            },
            removeEvent: function() {
                this.player.off("init"),
                this.player.off("ready"),
                this.player.off("ended"),
                this.player.off("play"),
                this.player.off("pause"),
                this.player.off("seekStart"),
                this.player.off("seekEnd"),
                this.player.off("canplaythrough"),
                this.player.off("error"),
                clearInterval(d)
            },
            _onPlayerInit: function() {
                this.sessionId = h.guid(),
                this._log("INIT", {}),
                this.buffer_flag = 0,
                this.pause_flag = 0
            },
            _onPlayerClose: function() {
                this._log("CLOSE", {
                    vt: Math.floor(1e3 * this.player.getCurrentTime())
                })
            },
            _onPlayerReady: function() {
                this.startTimePlay = (new Date).getTime()
            },
            _onPlayerFinish: function() {
                this.sessionId = h.guid(),
                this._log("STOP", {
                    vt: Math.floor(1e3 * this.player.getCurrentTime())
                })
            },
            _onPlayerPlay: function() {
                return ! this.buffer_flag && this.player._options.autoplay ? (this.first_play_time = (new Date).getTime(), this._log("PLAY", {
                    dsm: "fix",
                    vt: 0,
                    cost: this.first_play_time - this.player.getReadyTime()
                }), void(this.buffer_flag = 1)) : void(this.buffer_flag && this.pause_flag && (this.pause_flag = 0, this.pauseEndTime = (new Date).getTime(), this._log("RECOVER", {
                    vt: Math.floor(1e3 * this.player.getCurrentTime()),
                    cost: this.pauseEndTime - this.pauseTime
                })))
            },
            _onPlayerPause: function() {
                this.buffer_flag && this.startTimePlay && (this.seeking || (this.pause_flag = 1, this.pauseTime = (new Date).getTime(), this._log("PAUSE", {
                    vt: Math.floor(1e3 * this.player.getCurrentTime())
                })))
            },
            _onPlayerSeekStart: function(a) {
                this.seekStartTime = a.paramData.fromTime,
                this.seeking = !0,
                this.seekStartStamp = (new Date).getTime()
            },
            _onPlayerSeekEnd: function(a) {
                this.seekEndStamp = (new Date).getTime(),
                this._log("SEEK", {
                    drag_from_timestamp: Math.floor(1e3 * this.seekStartTime),
                    drag_to_timestamp: Math.floor(1e3 * a.paramData.toTime)
                }),
                this._log("SEEK_END", {
                    vt: Math.floor(1e3 * this.player.getCurrentTime()),
                    cost: this.seekEndStamp - this.seekStartStamp
                }),
                this.seeking = !1
            },
            _onPlayerLoaded: function() {
                this.buffer_flag && this.startTimePlay && (this.stucking || this.seeking || (this.stuckStartTime = (new Date).getTime(), this.stuckStartTime - this.startTimePlay <= 1e3 || (this.stucking = !0, this._log("UNDERLOAD", {
                    vt: Math.floor(1e3 * this.player.getCurrentTime())
                }), this.stuckStartTime = (new Date).getTime())))
            },
            _onPlayerUnderload: function() {
                if (!this.buffer_flag && !this.player._options.autoplay) return this.first_play_time = (new Date).getTime(),
                this._log("PLAY", {
                    play_mode: "fix",
                    vt: 0,
                    cost: this.first_play_time - this.player.getReadyTime()
                }),
                void(this.buffer_flag = 1);
                if ((this.buffer_flag || !this.player._options.autoplay) && this.stucking && !this.seeking) {
                    var a = Math.floor(1e3 * this.player.getCurrentTime()),
                    b = this.stuckStartTime || (new Date).getTime(),
                    c = Math.floor((new Date).getTime() - b);
                    c < 0 && (c = 0),
                    this._log("LOADED", {
                        vt: a,
                        cost: c
                    }),
                    this.stucking = !1
                }
            },
            _onPlayerHeartBeat: function() {
                if (!this.seeking) {
                    var a = Math.floor(1e3 * this.player.getCurrentTime()),
                    b = this;
                    this.timer || (this.timer = setTimeout(function() { ! b.seeking && b._log("HEARTBEAT", {
                            progress: a
                        }),
                        clearTimeout(b.timer),
                        b.timer = null
                    },
                    6e4))
                }
            },
            _onPlayerError: function() {
                var a, b = {
                    MEDIA_ERR_NETWORK: -1,
                    MEDIA_ERR_SRC_NOT_SUPPORTED: -2,
                    MEDIA_ERR_DECODE: -3
                },
                c = this.player.getError(),
                d = c.code;
                f.each(c.__proto__,
                function(b, c) {
                    if (c === d) return a = b,
                    !1
                }),
                b[a] && this._log("ERROR", {
                    vt: Math.floor(1e3 * this.player.getCurrentTime()),
                    error_code: b[a],
                    error_msg: a
                })
            },
            _log: function(a, b) {
                var c = f.copy(this.opt),
                d = k.logReportTo;
                c.e = l[a],
                c.s = this.sessionId,
                c.ct = (new Date).getTime();
                var e = [];
                f.each(b,
                function(a, b) {
                    e.push(a + "=" + b)
                }),
                e = e.join("&"),
                "" == e && (e = "0"),
                c.args = encodeURIComponent(e);
                var g = [];
                f.each(c,
                function(a, b) {
                    g.push(a + "=" + b)
                }),
                g = g.join("&"),
                i.jsonp(d + "?" + g,
                function() {},
                function() {})
            },
            _getUuid: function() {
                var a = g.get("p_h5_u");
                return a || (a = h.guid(), g.set("p_h5_u", a, 7)),
                a
            }
        });
        b.exports = m
    },
    {
        "../config": 1,
        "../lib/cookie": 3,
        "../lib/data": 4,
        "../lib/io": 8,
        "../lib/object": 10,
        "../lib/oo": 11,
        "../lib/ua": 12
    }],
    16 : [function(require, module, exports) {
        function sleep(a) {
            for (var b = Date.now(); Date.now() - b <= a;);
        }
        var Component = require("../ui/component"),
        _ = require("../lib/object"),
        Dom = require("../lib/dom"),
        Event = require("../lib/event"),
        io = require("../lib/io"),
        UI = require("../ui/exports"),
        Monitor = require("../monitor/monitor"),
        UA = require("../lib/ua"),
        debug_flag = 0,
        Player = Component.extend({
            init: function(tag, options) {
                if (this.tag = tag, this.loaded = !1, this.played = !1, Component.call(this, this, options), options.plugins && _.each(options.plugins,
                function(a, b) {
                    this[a](b)
                },
                this), options.useNativeControls ? this.tag.setAttribute("controls", "controls") : (this.UI = UI, this.initChildren()), this.bindVideoEvent(), this._options.source ? (this._options.trackLog && (this._monitor = new Monitor(this, {
                    video_id: 0,
                    album_id: 0,
                    from: this._options.from
                })), this.trigger("init"), debug_flag && console.log("init"), (this._options.autoplay || this._options.preload) && (this.getMetaData(), this.tag.setAttribute("src", this._options.source), this.readyTime = (new Date).getTime(), this.loaded = !0)) : this._options.vid ? this.loadVideoInfo() : (this._options.trackLog && (this._monitor = new Monitor(this, {
                    video_id: 0,
                    album_id: 0,
                    from: this._options.from
                })), this.trigger("init"), debug_flag && console.log("init")), this._options.extraInfo) {
                    var dict = eval(this._options.extraInfo);
                    dict.liveRetry && (this._options.liveRetry = dict.liveRetry)
                }
                this.on("readyState",
                function() {
                    this.trigger("ready"),
                    debug_flag && console.log("ready")
                })
            }
        });
        Player.prototype.initChildren = function() {
            var a = this.options(),
            b = a.skinLayout;
            if (b !== !1 && !_.isArray(b)) throw new Error("PrismPlayer Error: skinLayout should be false or type of array!");
            b !== !1 && 0 !== b.length && (this.options({
                children: b
            }), Component.prototype.initChildren.call(this)),
            this.trigger("uiH5Ready"),
            debug_flag && console.log("uiH5ready")
        },
        Player.prototype.createEl = function() {
            "VIDEO" !== this.tag.tagName && (this._el = this.tag, this.tag = Component.prototype.createEl.call(this, "video"), this._options.playsinline && (this.tag.setAttribute("webkit-playsinline", ""), this.tag.setAttribute("playsinline", "")));
            var a = this._el,
            b = this.tag;
            b.player = this;
            var c = Dom.getElementAttributes(b);
            return _.each(c,
            function(b) {
                a.setAttribute(b, c[b])
            }),
            this.setVideoAttrs(),
            b.parentNode && b.parentNode.insertBefore(a, b),
            Dom.insertFirst(b, a),
            this.cover = Dom.createEl("div"),
            Dom.addClass(this.cover, "prism-cover"),
            a.appendChild(this.cover),
            this.options().cover && (this.cover.style.backgroundImage = "url(" + this.options().cover + ")"),
            UA.IS_IOS && Dom.css(b, "display", "none"),
            a
        },
        Player.prototype.setVideoAttrs = function() {
            var a = this._options.preload,
            b = this._options.autoplay;

            if(navigator.userAgent.toLowerCase().indexOf('iphone') != -1){
                this.tag.setAttribute("controls", "controls")
            }

            this.tag.style.width = "100%",
            this.tag.style.height = "100%",
            this.tag.setAttribute("id", "livevideo"),
            this.tag.setAttribute("x5-video-player-type", "h5"),
            this.tag.setAttribute("x5-video-player-fullscreen", "true"),
            this.tag.setAttribute("x5-video-orientation", "portrait"),
            this.tag.setAttribute("webkit-playsinline", ""),
            this.tag.setAttribute("playsinline", ""),
            a && this.tag.setAttribute("preload", "preload"),
            b && this.tag.setAttribute("autoplay", "autoplay")
				
        },
        Player.prototype.id = function() {
            return this.el().id
        },
        Player.prototype.renderUI = function() {},
        Player.prototype.bindVideoEvent = function() {
            var a = this.tag,
            b = this;
            Event.on(a, "loadstart",
            function(a) {
                b.trigger("loadstart"),
                debug_flag && console.log("loadstart")
            }),
            Event.on(a, "durationchange",
            function(a) {
                b.trigger("durationchange"),
                debug_flag && console.log("durationchange")
            }),
            Event.on(a, "loadedmetadata",
            function(a) {
                b.trigger("loadedmetadata"),
                debug_flag && console.log("loadedmetadata")
            }),
            Event.on(a, "loadeddata",
            function(a) {
                b.trigger("loadeddata"),
                debug_flag && console.log("loadeddata")
            }),
            Event.on(a, "progress",
            function(a) {
                b.trigger("progress"),
                debug_flag && console.log("progress")
            }),
            Event.on(a, "canplay",
            function(a) {
                var c = (new Date).getTime() - b.readyTime;
                b.trigger("canplay", {
                    loadtime: c
                }),
                debug_flag && console.log("canplay")
            }),
            Event.on(a, "canplaythrough",
            function(c) {
                b.cover && b._options.autoplay && (Dom.css(b.cover, "display", "none"), delete b.cover),
                "none" === a.style.display && UA.IS_IOS && setTimeout(function() {
                    Dom.css(a, "display", "block")
                },
                100),
                b.trigger("canplaythrough"),
                debug_flag && console.log("canplaythrough")
            }),
            Event.on(a, "play",
            function(a) {
                b.trigger("play"),
                debug_flag && console.log("play")
            }),
            Event.on(a, "play",
            function(a) {
                b.trigger("videoRender"),
                debug_flag && console.log("videoRender")
            }),
            Event.on(a, "pause",
            function(a) {
                b.trigger("pause"),
                debug_flag && console.log("pause")
            }),
            Event.on(a, "ended",
            function(a) {
                b._options.rePlay && (b.seek(0), b.tag.play()),
                b.trigger("ended"),
                debug_flag && console.log("ended")
            }),
            Event.on(a, "stalled",
            function(a) {
                b.trigger("stalled"),
                debug_flag && console.log("stalled")
            }),
            Event.on(a, "waiting",
            function(a) {
                b.trigger("waiting"),
                debug_flag && console.log("waiting")
            }),
            Event.on(a, "playing",
            function(a) {
                b.trigger("playing"),
                debug_flag && console.log("playing")
            }),
            Event.on(a, "error",
            function(a) {
                if (console.log("error"), b._options.isLive) b._options.liveRetry ? (sleep(2e3), b.tag.load(b._options.source), b.tag.play()) : b.trigger("error"),
                b.trigger("liveStreamStop");
                else {
                    var c = 0;
                    b._options.source.indexOf("flv") > 0 ? c = 1 : b._options.source.indexOf("m3u8") > 0 && !UA.IS_MOBILE && (c = 1),
                    errmsg = document.querySelector("#" + b.id()),
                    errmsg.style.lineHeight = errmsg.clientHeight + "px",
                    Dom.css(errmsg, "text-align", "center"),
                    Dom.css(errmsg, "color", "#FFFFFF"),
                    c ? errmsg.innerText = "播放失败: h5不支持此格式，请安装flashplayer": errmsg.innerText = "播放失败: 请确认播放源",
                    b.trigger("error")
                }
            }),
            Event.on(a, "onM3u8Retry",
            function(a) {
                b.trigger("m3u8Retry"),
                debug_flag && console.log("m3u8Retry")
            }),
            Event.on(a, "liveStreamStop",
            function(a) {
                b.trigger("liveStreamStop"),
                debug_flag && console.log("liveStreamStop")
            }),
            Event.on(a, "seeking",
            function(a) {
                b.trigger("seeking"),
                debug_flag && console.log("seeking")
            }),
            Event.on(a, "seeked",
            function(a) {
                b.trigger("seeked"),
                debug_flag && console.log("seeked")
            }),
            Event.on(a, "ratechange",
            function(a) {
                b.trigger("ratechange"),
                debug_flag && console.log("ratechange")
            }),
            Event.on(a, "timeupdate",
            function(a) {
                b.trigger("timeupdate"),
                debug_flag && console.log("timeupdate")
            }),
            Event.on(a, "webkitfullscreenchange",
            function(a) {
                b.trigger("fullscreenchange"),
                debug_flag && console.log("fullscreenchange")
            }),
            this.on("requestFullScreen",
            function() {
                Dom.addClass(b.el(), "prism-fullscreen"),
                debug_flag && console.log("request-fullscreen")
            }),
            this.on("cancelFullScreen",
            function() {
                Dom.removeClass(b.el(), "prism-fullscreen"),
                debug_flag && console.log("cancel-fullscreen")
            }),
            Event.on(a, "suspend",
            function(a) {
                b.trigger("suspend"),
                debug_flag && console.log("sudpend")
            }),
            Event.on(a, "abort",
            function(a) {
                b.trigger("abort"),
                debug_flag && console.log("abort")
            }),
            Event.on(a, "volumechange",
            function(a) {
                b.trigger("volumechange"),
                debug_flag && console.log("volumechange")
            }),
            Event.on(a, "drag",
            function(a) {
                b.trigger("drag"),
                debug_flag && console.log("drag")
            }),
            Event.on(a, "dragstart",
            function(a) {
                b.trigger("dragstart"),
                debug_flag && console.log("dragstart")
            }),
            Event.on(a, "dragover",
            function(a) {
                b.trigger("dragover"),
                debug_flag && console.log("dragover")
            }),
            Event.on(a, "dragenter",
            function(a) {
                b.trigger("dragenter"),
                debug_flag && console.log("dragenter")
            }),
            Event.on(a, "dragleave",
            function(a) {
                b.trigger("dragleave"),
                debug_flag && console.log("dragleave")
            }),
            Event.on(a, "ondrag",
            function(a) {
                b.trigger("ondrag"),
                debug_flag && console.log("ondrag")
            }),
            Event.on(a, "ondragstart",
            function(a) {
                b.trigger("ondragstart"),
                debug_flag && console.log("ondragstart")
            }),
            Event.on(a, "ondragover",
            function(a) {
                b.trigger("ondragover"),
                debug_flag && console.log("ondragover")
            }),
            Event.on(a, "ondragenter",
            function(a) {
                b.trigger("ondragenter"),
                debug_flag && console.log("ondragenter")
            }),
            Event.on(a, "ondragleave",
            function(a) {
                b.trigger("ondragleave"),
                debug_flag && console.log("ondragleave")
            }),
            Event.on(a, "drop",
            function(a) {
                b.trigger("drop"),
                debug_flag && console.log("drop")
            }),
            Event.on(a, "dragend",
            function(a) {
                b.trigger("dragend"),
                debug_flag && console.log("dragend")
            }),
            Event.on(a, "onscroll",
            function(a) {
                b.trigger("onscroll"),
                debug_flag && console.log("onscroll")
            })
        },
        Player.prototype.loadVideoInfo = function() {
            var a = this._options.vid,
            b = this;
            if (!a) throw new Error("PrismPlayer Error: vid should not be null!");
            io.jsonp("//tv.taobao.com/player/json/getBaseVideoInfo.do?vid=" + a + "&playerType=3",
            function(c) {
                if (1 !== c.status || !c.data.source) throw new Error("PrismPlayer Error: #vid:" + a + " cannot find video resource!");
                var d, e = -1;
                _.each(c.data.source,
                function(a, b) {
                    var c = +a.substring(1);
                    c > e && (e = c)
                }),
                d = c.data.source["v" + e],
                d = _.unescape(d),
                b._options.source = d,
                b._options.trackLog && (b._monitor = new Monitor(b, {
                    video_id: a,
                    album_id: c.data.baseInfo.aid,
                    from: b._options.from
                })),
                b.trigger("init"),
                debug_flag && console.log("init"),
                (b._options.autoplay || b._options.preload) && (b.getMetaData(), b.tag.setAttribute("src", b._options.source), b.readyTime = (new Date).getTime(), b.loaded = !0)
            },
            function() {
                throw new Error("PrismPlayer Error: network error!")
            })
        },
        Player.prototype.setControls = function() {
            var a = this.options();
            if (a.useNativeControls) this.tag.setAttribute("controls", "controls");
            else if ("object" == typeof a.controls) {
                var b = this._initControlBar(a.controls);
                this.addChild(b)
            }
        },
        Player.prototype._initControlBar = function(a) {
            var b = new ControlBar(this, a);
            return b
        },
        Player.prototype.getMetaData = function() {
            var a = this,
            b = null,
            c = this.tag;
            b = window.setInterval(function(d) {
                if (c.readyState > 0) {
                    var e = Math.round(c.duration);
                    a.tag.duration = e,
                    a.trigger("readyState"),
                    debug_flag && console.log("readystate"),
                    clearInterval(b)
                }
            },
            100)
        },
        Player.prototype.getReadyTime = function() {
            return this.readyTime
        },
        Player.prototype.readyState = function() {
            return this.tag.readyState
        },
        Player.prototype.getError = function() {
            return this.tag.error
        },
        Player.prototype.play = function() {
            var a = this;
            return this._options.autoplay || this._options.preload || this.loaded || (this.getMetaData(), this.tag.setAttribute("src", this._options.source), this.readyTime = (new Date).getTime(), this.loaded = !0),
            a.cover && !a._options.autoplay && (Dom.css(a.cover, "display", "none"), delete a.cover),
            this.tag.play(),
            this
        },
        Player.prototype.replay = function() {
            return this.seek(0),
            this.tag.play(),
            this
        },
        Player.prototype.pause = function() {
            return this.tag.pause(),
            this
        },
        Player.prototype.stop = function() {
            return this.tag.setAttribute("src", null),
            this
        },
        Player.prototype.paused = function() {
            return this.tag.paused !== !1
        },
        Player.prototype.getDuration = function() {
            var a = this.tag.duration;
            return a
        },
        Player.prototype.getCurrentTime = function() {
            var a = this.tag.currentTime;
            return a
        },
        Player.prototype.seek = function(a) {
            a === this.tag.duration && a--;
            try {
                this.tag.currentTime = a
            } catch(a) {
                console.log(a)
            }
            return this
        },
        Player.prototype.loadByVid = function(a) {
            this._options.vid = a;
            var b = this;
            if (!a) throw new Error("PrismPlayer Error: vid should not be null!");
            io.jsonp("//tv.taobao.com/player/json/getBaseVideoInfo.do?vid=" + a + "&playerType=3",
            function(c) {
                if (1 !== c.status || !c.data.source) throw new Error("PrismPlayer Error: #vid:" + a + " cannot find video resource!");
                var d, e = -1;
                _.each(c.data.source,
                function(a, b) {
                    var c = +a.substring(1);
                    c > e && (e = c)
                }),
                d = c.data.source["v" + e],
                d = _.unescape(d),
                b._options.source = d,
                b._options.trackLog && (b._monitor ? b._monitor.updateVideoInfo({
                    video_id: a,
                    album_id: c.data.baseInfo.aid,
                    from: b._options.from
                }) : b._monitor = new Monitor(b, {
                    video_id: a,
                    album_id: c.data.baseInfo.aid,
                    from: b._options.from
                })),
                b._options.autoplay = !0,
                b.loaded || (b.trigger("init"), debug_flag && console.log("init")),
                b.getMetaData(),
                b.tag.setAttribute("src", b._options.source),
                b.readyTime = (new Date).getTime(),
                b.loaded = !0,
                b.cover && b._options.autoplay && (Dom.css(b.cover, "display", "none"), delete b.cover),
                b.tag.play()
            },
            function() {
                throw new Error("PrismPlayer Error: network error!")
            })
        },
        Player.prototype.loadByUrl = function(a, b) {
            this._options.vid = 0,
            this._options.source = a,
            this._options.autoplay = !0,
            this._options.trackLog && (this._monitor ? this._monitor.updateVideoInfo({
                video_id: 0,
                album_id: 0,
                from: this._options.from
            }) : this._monitor = new Monitor(this, {
                video_id: 0,
                album_id: 0,
                from: this._options.from
            })),
            this.loaded || (this.trigger("init"), debug_flag && console.log("init")),
            this.getMetaData(),
            this.tag.setAttribute("src", this._options.source),
            this.readyTime = (new Date).getTime(),
            this.loaded = !0,
            this.cover && (this._options.preload || this._options.autoplay) && (Dom.css(this.cover, "display", "none"), delete this.cover),
            this.tag.play(),
            b && !isNaN(b) && this.seek(b)
        },
        Player.prototype.dispose = function() {
            this.tag.pause();
            var a = this.tag;
            Event.off(a, "timeupdate"),
            Event.off(a, "play"),
            Event.off(a, "pause"),
            Event.off(a, "canplay"),
            Event.off(a, "waiting"),
            Event.off(a, "playing"),
            Event.off(a, "ended"),
            Event.off(a, "error"),
            Event.off(a, "durationchange"),
            Event.off(a, "loadedmetadata"),
            Event.off(a, "loadeddata"),
            Event.off(a, "progress"),
            Event.off(a, "canplaythrough"),
            Event.off(a, "webkitfullscreenchange"),
            this.tag = null,
            this._options = null,
            this._monitor && (this._monitor.removeEvent(), this._monitor = null)
        },
        Player.prototype.mute = function() {
            return this.tag.muted = !0,
            this
        },
        Player.prototype.unMute = function() {
            return this.tag.muted = !1,
            this
        },
        Player.prototype.muted = function() {
            return this.tag.muted
        },
        Player.prototype.getVolume = function() {
            return this.tag.volume
        },
        Player.prototype.getOptions = function() {
            return this._options
        },
        Player.prototype.setVolume = function(a) {
            this.tag.volume = a
        },
        Player.prototype.hideProgress = function() {
            var a = this;
            a.trigger("hideProgress")
        },
        Player.prototype.cancelHideProgress = function() {
            var a = this;
            a.trigger("cancelHideProgress")
        },
        Player.prototype.setPlayerSize = function(a, b) {
            if (this._el.style.width = a, b) {
                var c = b.indexOf("%");
                if (c > 0) {
                    var d = window.screen.height,
                    e = b.replace("%", "");
                    if (isNaN(e)) this._el.style.height = b;
                    else {
                        var f = 9 * d * parseInt(e) / 1e3;
                        this._el.style.height = String(f % 2 ? f + 1 : f) + "px"
                    }
                } else this._el.style.height = b
            }
        };
        var __supportFullscreen = function() {
            var a, b;
            b = Dom.createEl("div"),
            a = {};
            var c = [["requestFullscreen", "exitFullscreen", "fullscreenElement", "fullscreenEnabled", "fullscreenchange", "fullscreenerror", "fullScreen"], ["webkitRequestFullscreen", "webkitExitFullscreen", "webkitFullscreenElement", "webkitFullscreenEnabled", "webkitfullscreenchange", "webkitfullscreenerror", "webkitfullScreen"], ["webkitRequestFullScreen", "webkitCancelFullScreen", "webkitCurrentFullScreenElement", "webkitFullscreenEnabled", "webkitfullscreenchange", "webkitfullscreenerror", "webkitIsFullScreen"], ["mozRequestFullScreen", "mozCancelFullScreen", "mozFullScreenElement", "mozFullScreenEnabled", "mozfullscreenchange", "mozfullscreenerror", "mozfullScreen"], ["msRequestFullscreen", "msExitFullscreen", "msFullscreenElement", "msFullscreenEnabled", "MSFullscreenChange", "MSFullscreenError", "MSFullScreen"]];
            if (UA.IS_IOS) a.requestFn = "webkitEnterFullscreen",
            a.cancelFn = "webkitExitFullscreen",
            a.eventName = "webkitfullscreenchange",
            a.isFullScreen = "webkitDisplayingFullscreen";
            else {
                for (var d = 5,
                e = 0; e < d; e++) if (c[e][1] in document) {
                    a.requestFn = c[e][0],
                    a.cancelFn = c[e][1],
                    a.eventName = c[e][4],
                    a.isFullScreen = c[e][6];
                    break
                }
                "requestFullscreen" in document ? a.requestFn = "requestFullscreen": "webkitRequestFullscreen" in document ? a.requestFn = "webkitRequestFullscreen": "webkitRequestFullScreen" in document ? a.requestFn = "webkitRequestFullScreen": "webkitEnterFullscreen" in document ? a.requestFn = "webkitEnterFullscreen": "mozRequestFullScreen" in document ? a.requestFn = "mozRequestFullScreen": "msRequestFullscreen" in document && (a.requestFn = "msRequestFullscreen"),
                "fullscreenchange" in document ? a.eventName = "fullscreenchange": "webkitfullscreenchange" in document ? a.eventName = "webkitfullscreenchange": "webkitfullscreenchange" in document ? a.eventName = "webkitfullscreenchange": "webkitfullscreenchange" in document ? a.eventName = "webkitfullscreenchange": "mozfullscreenchange" in document ? a.eventName = "mozfullscreenchange": "MSFullscreenChange" in document && (a.eventName = "MSFullscreenChange"),
                "fullScreen" in document ? a.isFullScreen = "fullScreen": "webkitfullScreen" in document ? a.isFullScreen = "webkitfullScreen": "webkitIsFullScreen" in document ? a.isFullScreen = "webkitIsFullScreen": "webkitDisplayingFullscreen" in document ? a.isFullScreen = "webkitDisplayingFullscreen": "mozfullScreen" in document ? a.isFullScreen = "mozfullScreen": "MSFullScreen" in document && (a.isFullScreen = "MSFullScreen")
            }
            return a.requestFn ? a: null
        } ();
        Player.prototype._enterFullWindow = function() {
            this.isFullWindow = !0,
            this.docOrigOverflow = document.documentElement.style.overflow,
            document.documentElement.style.overflow = "hidden",
            Dom.addClass(document.getElementsByTagName("body")[0], "prism-full-window")
        },
        Player.prototype._exitFullWindow = function() {
            this.isFullWindow = !1,
            document.documentElement.style.overflow = this.docOrigOverflow,
            Dom.removeClass(document.getElementsByTagName("body")[0], "prism-full-window")
        },
        Player.prototype.requestFullScreen = function() {
            var a = __supportFullscreen,
            b = this.el(),
            c = this;
            return UA.IS_IOS ? (b = this.tag, b[a.requestFn](), this) : (this.isFullScreen = !0, a ? (Event.on(document, a.eventName,
            function(b) {
                c.isFullScreen = document[a.isFullScreen],
                c.isFullScreen === !0 && Event.off(document, a.eventName),
                c.trigger("requestFullScreen")
            }), b[a.requestFn]()) : (this._enterFullWindow(), this.trigger("requestFullScreen")), this)
        },
        Player.prototype.cancelFullScreen = function() {
            var a = __supportFullscreen,
            b = this;
            return this.isFullScreen = !1,
            a ? (Event.on(document, a.eventName,
            function(c) {
                b.isFullScreen = document[a.isFullScreen],
                b.isFullScreen === !1 && Event.off(document, a.eventName),
                b.trigger("cancelFullScreen")
            }), document[a.cancelFn](), this.trigger("play")) : (this._exitFullWindow(), this.trigger("cancelFullScreen"), this.trigger("play")),
            this
        },
        Player.prototype.getIsFullScreen = function() {
            return this.isFullScreen
        },
        Player.prototype.getBuffered = function() {
            return this.tag.buffered
        },
        Player.prototype.setToastEnabled = function(a) {},
        Player.prototype.setLoadingInvisible = function() {},
        module.exports = Player
    },
    {
        "../lib/dom": 5,
        "../lib/event": 6,
        "../lib/io": 8,
        "../lib/object": 10,
        "../lib/ua": 12,
        "../monitor/monitor": 15,
        "../ui/component": 17,
        "../ui/exports": 26
    }],
    17 : [function(a, b, c) {
        var d = a("../lib/oo"),
        e = a("../lib/data"),
        f = a("../lib/object"),
        g = a("../lib/dom"),
        h = a("../lib/event"),
        i = a("../lib/function"),
        j = a("../lib/layout"),
        k = d.extend({
            init: function(a, b) {
                var c = this;
                this._player = a,
                this._options = f.copy(b),
                this._el = this.createEl(),
                this._id = a.id() + "_component_" + e.guid(),
                this._children = [],
                this._childIndex = {},
                this._player.on("uiH5Ready",
                function() {
                    c.renderUI(),
                    c.syncUI(),
                    c.bindEvent()
                })
            }
        });
        k.prototype.renderUI = function() {
            j.render(this.el(), this.options()),
            this.el().id = this.id()
        },
        k.prototype.syncUI = function() {},
        k.prototype.bindEvent = function() {},
        k.prototype.createEl = function(a, b) {
            return g.createEl(a, b)
        },
        k.prototype.options = function(a) {
            return void 0 === a ? this._options: this._options = f.merge(this._options, a)
        },
        k.prototype.el = function() {
            return this._el
        },
        k.prototype._contentEl,
        k.prototype.player = function() {
            return this._player
        },
        k.prototype.contentEl = function() {
            return this._contentEl || this._el
        },
        k.prototype._id,
        k.prototype.id = function() {
            return this._id
        },
        k.prototype.addChild = function(a, b) {
            var c;
            if ("string" == typeof a) {
                if (!this._player.UI[a]) return;
                c = new this._player.UI[a](this._player, b)
            } else c = a;
            return this._children.push(c),
            "function" == typeof c.id && (this._childIndex[c.id()] = c),
            "function" == typeof c.el && c.el() && this.contentEl().appendChild(c.el()),
            c
        },
        k.prototype.removeChild = function(a) {
            if (a && this._children) {
                for (var b = !1,
                c = this._children.length - 1; c >= 0; c--) if (this._children[c] === a) {
                    b = !0,
                    this._children.splice(c, 1);
                    break
                }
                if (b) {
                    this._childIndex[a.id] = null;
                    var d = a.el();
                    d && d.parentNode === this.contentEl() && this.contentEl().removeChild(a.el())
                }
            }
        },
        k.prototype.initChildren = function() {
            var a, b, c, d, e;
            if (a = this, b = this.options().children) if (f.isArray(b)) for (var g = 0; g < b.length; g++) c = b[g],
            "string" == typeof c ? (d = c, e = {}) : (d = c.name, e = c),
            a.addChild(d, e);
            else f.each(b,
            function(b, c) {
                c !== !1 && a.addChild(b, c)
            })
        },
        k.prototype.on = function(a, b) {
            return h.on(this._el, a, i.bind(this, b)),
            this
        },
        k.prototype.off = function(a, b) {
            return h.off(this._el, a, b),
            this
        },
        k.prototype.one = function(a, b) {
            return h.one(this._el, a, i.bind(this, b)),
            this
        },
        k.prototype.trigger = function(a, b) {
            return b && (this._el.paramData = b),
            h.trigger(this._el, a),
            this
        },
        k.prototype.addClass = function(a) {
            return g.addClass(this._el, a),
            this
        },
        k.prototype.removeClass = function(a) {
            return g.removeClass(this._el, a),
            this
        },
        k.prototype.show = function() {
            return this._el.style.display = "block",
            this
        },
        k.prototype.hide = function() {
            return this._el.style.display = "none",
            this
        },
        k.prototype.destroy = function() {
            if (this.trigger({
                type: "destroy",
                bubbles: !1
            }), this._children) for (var a = this._children.length - 1; a >= 0; a--) this._children[a].destroy && this._children[a].destroy();
            this.children_ = null,
            this.childIndex_ = null,
            this.off(),
            this._el.parentNode && this._el.parentNode.removeChild(this._el),
            e.removeData(this._el),
            this._el = null
        },
        b.exports = k
    },
    {
        "../lib/data": 4,
        "../lib/dom": 5,
        "../lib/event": 6,
        "../lib/function": 7,
        "../lib/layout": 9,
        "../lib/object": 10,
        "../lib/oo": 11
    }],
    18 : [function(a, b, c) {
        var d = a("../component"),
        e = a("../../lib/dom"),
        f = d.extend({
            init: function(a, b) {
                d.call(this, a, b),
                this.addClass(b.className || "prism-big-play-btn")
            },
            bindEvent: function() {
                var a = this;
                this._player.on("play",
                function() {
                    a.addClass("playing"),
                    e.css(a.el(), "display", "none")
                }),
                this._player.on("pause",
                function() {
                    a.removeClass("playing"),
                    e.css(a.el(), "display", "block")
                }),
                this.on("click",
                function() {
                    a._player.paused() && (a._player.play(), e.css(a.el(), "display", "none"))
                })
            }
        });
        b.exports = f
    },
    {
        "../../lib/dom": 5,
        "../component": 17
    }],
    19 : [function(a, b, c) {
        var d = a("../component"),
        e = d.extend({
            init: function(a, b) {
                d.call(this, a, b),
                this.addClass(b.className || "prism-controlbar"),
                this.initChildren(),
                this.onEvent()
            },
            createEl: function() {
                var a = d.prototype.createEl.call(this);
                return a.innerHTML = '<div class="prism-controlbar-bg"></div>',
                a
            },
            onEvent: function() {
                var a = this.player(),
                b = this;
                this.timer = null,
                a.on("click",
                function(a) {
                    a.preventDefault(),
                    a.stopPropagation(),
                    b._show(),
                    b._hide()
                }),
                a.on("ready",
                function() {
                    b._hide()
                }),
                this.on("touchstart",
                function() {
                    b._show()
                }),
                this.on("touchmove",
                function() {
                    b._show()
                }),
                this.on("touchend",
                function() {
                    b._hide()
                })
            },
            _show: function() {
                this.show(),
                this._player.trigger("showBar"),
                this.timer && (clearTimeout(this.timer), this.timer = null)
            },
            _hide: function() {
                var a = this,
                b = this.player(),
                c = b.options(),
                d = c.showBarTime;
                this.timer = setTimeout(function() {
                    a.hide(),
                    a._player.trigger("hideBar")
                },
                d)
            }
        });
        b.exports = e
    },
    {
        "../component": 17
    }],
    20 : [function(a, b, c) {
        var d = a("../component"),
        e = d.extend({
            init: function(a, b) {
                d.call(this, a, b),
                this.addClass(b.className || "prism-fullscreen-btn")
            },
            bindEvent: function() {
                var a = this;
                this._player.on("requestFullScreen",
                function() {
                    a.addClass("fullscreen")
                }),
                this._player.on("cancelFullScreen",
                function() {
                    a.removeClass("fullscreen")
                }),
                this.on("click",
                function() {
                    this._player.getIsFullScreen() ? this._player.cancelFullScreen() : this._player.requestFullScreen()
                })
            }
        });
        b.exports = e
    },
    {
        "../component": 17
    }],
    21 : [function(a, b, c) {
        var d = a("../component"),
        e = (a("../../lib/util"), d.extend({
            init: function(a, b) {
                d.call(this, a, b),
                this.className = b.className ? b.className: "prism-live-display",
                this.addClass(this.className)
            }
        }));
        b.exports = e
    },
    {
        "../../lib/util": 14,
        "../component": 17
    }],
    22 : [function(a, b, c) {
        var d = a("../component"),
        e = d.extend({
            init: function(a, b) {
                d.call(this, a, b),
                this.addClass(b.className || "prism-play-btn")
            },
            bindEvent: function() {
                var a = this;
                this._player.on("play",
                function() {
                    a.addClass("playing")
                }),
                this._player.on("pause",
                function() {
                    a.removeClass("playing")
                }),
                this.on("click",
                function() {
                    a._player.paused() ? (a._player.play(), a.addClass("playing")) : (a._player.pause(), a.removeClass("playing"))
                })
            }
        });
        b.exports = e
    },
    {
        "../component": 17
    }],
    23 : [function(a, b, c) {
        var d = a("../component"),
        e = a("../../lib/dom"),
        f = a("../../lib/event"),
        g = a("../../lib/ua"),
        h = a("../../lib/function"),
        i = d.extend({
            init: function(a, b) {
                d.call(this, a, b),
                this.className = b.className ? b.className: "prism-progress",
                this.addClass(this.className)
            },
            createEl: function() {
                var a = d.prototype.createEl.call(this);
                return a.innerHTML = '<div class="prism-progress-loaded"></div><div class="prism-progress-played"></div><div class="prism-progress-cursor"></div>',
                a
            },
            bindEvent: function() {
                var a = this;
                this.loadedNode = document.querySelector("#" + this.id() + " .prism-progress-loaded"),
                this.playedNode = document.querySelector("#" + this.id() + " .prism-progress-played"),
                this.cursorNode = document.querySelector("#" + this.id() + " .prism-progress-cursor"),
                this.controlNode = document.getElementsByClassName("prism-controlbar")[0],
                f.on(this.cursorNode, "mousedown",
                function(b) {
                    a._onMouseDown(b)
                }),
                f.on(this.cursorNode, "touchstart",
                function(b) {
                    a._onMouseDown(b)
                }),
                f.on(this._el, "click",
                function(b) {
                    a._onMouseClick(b)
                }),
                this._player.on("hideProgress",
                function(b) {
                    a._hideProgress(b)
                }),
                this._player.on("cancelHideProgress",
                function(b) {
                    a._cancelHideProgress(b)
                }),
                this.bindTimeupdate = h.bind(this, this._onTimeupdate),
                this._player.on("timeupdate", this.bindTimeupdate),
                g.IS_IPAD ? this.interval = setInterval(function() {
                    a._onProgress()
                },
                500) : this._player.on("progress",
                function() {
                    a._onProgress()
                })
            },
            _hideProgress: function(a) {
                f.off(this.cursorNode, "mousedown"),
                f.off(this.cursorNode, "touchstart")
            },
            _cancelHideProgress: function(a) {
                var b = this;
                f.on(this.cursorNode, "mousedown",
                function(a) {
                    b._onMouseDown(a)
                }),
                f.on(this.cursorNode, "touchstart",
                function(a) {
                    b._onMouseDown(a)
                })
            },
            _onMouseClick: function(a) {
                var b = a.touches ? a.touches[0].pageX: a.pageX,
                c = b - this.el().offsetLeft,
                d = this.el().offsetWidth,
                e = this._player.getDuration() ? c / d * this._player.getDuration() : 0;
                e < 0 && (e = 0),
                e > this._player.getDuration() && (e = this._player.getDuration()),
                this._player.trigger("seekStart", {
                    fromTime: this._player.getCurrentTime()
                }),
                this._player.seek(e),
                this._player.play(),
                this._player.trigger("seekEnd", {
                    toTime: this._player.getCurrentTime()
                })
            },
            _onMouseDown: function(a) {
                var b = this;
                a.preventDefault(),
                this._player.pause(),
                this._player.trigger("seekStart", {
                    fromTime: this._player.getCurrentTime()
                }),
                f.on(this.controlNode, "mousemove",
                function(a) {
                    b._onMouseMove(a)
                }),
                f.on(this.controlNode, "touchmove",
                function(a) {
                    b._onMouseMove(a)
                }),
                f.on(this._player.tag, "mouseup",
                function(a) {
                    b._onPlayerMouseUp(a)
                }),
                f.on(this._player.tag, "touchend",
                function(a) {
                    b._onPlayerMouseUp(a)
                }),
                f.on(this.controlNode, "mouseup",
                function(a) {
                    b._onControlBarMouseUp(a)
                }),
                f.on(this.controlNode, "touchend",
                function(a) {
                    b._onControlBarMouseUp(a)
                })
            },
            _onMouseUp: function(a) {
                a.preventDefault(),
                f.off(this.controlNode, "mousemove"),
                f.off(this.controlNode, "touchmove"),
                f.off(this._player.tag, "mouseup"),
                f.off(this._player.tag, "touchend"),
                f.off(this.controlNode, "mouseup"),
                f.off(this.controlNode, "touchend");
                var b = this.playedNode.offsetWidth / this.el().offsetWidth * this._player.getDuration();
                this._player.getDuration();
                this._player.seek(b),
                this._player.play(),
                this._player.trigger("seekEnd", {
                    toTime: this._player.getCurrentTime()
                })
            },
            _onControlBarMouseUp: function(a) {
                a.preventDefault(),
                f.off(this.controlNode, "mousemove"),
                f.off(this.controlNode, "touchmove"),
                f.off(this._player.tag, "mouseup"),
                f.off(this._player.tag, "touchend"),
                f.off(this.controlNode, "mouseup"),
                f.off(this.controlNode, "touchend");
                var b = this.playedNode.offsetWidth / this.el().offsetWidth * this._player.getDuration();
                this._player.getDuration();
                this._player.seek(b),
                this._player.play(),
                this._player.trigger("seekEnd", {
                    toTime: this._player.getCurrentTime()
                })
            },
            _onPlayerMouseUp: function(a) {
                a.preventDefault(),
                f.off(this.controlNode, "mousemove"),
                f.off(this.controlNode, "touchmove"),
                f.off(this._player.tag, "mouseup"),
                f.off(this._player.tag, "touchend"),
                f.off(this.controlNode, "mouseup"),
                f.off(this.controlNode, "touchend");
                var b = this.playedNode.offsetWidth / this.el().offsetWidth * this._player.getDuration();
                this._player.getDuration();
                isNaN(b) || (this._player.seek(b), this._player.play()),
                this._player.trigger("seekEnd", {
                    toTime: this._player.getCurrentTime()
                })
            },
            _onMouseMove: function(a) {
                a.preventDefault();
                var b = a.touches ? a.touches[0].pageX: a.pageX,
                c = b - this.el().offsetLeft,
                d = this.el().offsetWidth,
                e = this._player.getDuration() ? c / d * this._player.getDuration() : 0;
                e < 0 && (e = 0),
                e > this._player.getDuration() && (e = this._player.getDuration()),
                this._player.seek(e),
                this._player.play(),
                this._updateProgressBar(this.playedNode, e),
                this._updateCursorPosition(e)
            },
            _onTimeupdate: function(a) {
                this._updateProgressBar(this.playedNode, this._player.getCurrentTime()),
                this._updateCursorPosition(this._player.getCurrentTime()),
                this._player.trigger("updateProgressBar", {
                    time: this._player.getCurrentTime()
                })
            },
            _onProgress: function(a) {
                this._player.getDuration() && this._player.getBuffered().length >= 1 && this._updateProgressBar(this.loadedNode, this._player.getBuffered().end(this._player.getBuffered().length - 1))
            },
            _updateProgressBar: function(a, b) {
                var c = this._player.getDuration() ? b / this._player.getDuration() : 0;
                a && e.css(a, "width", 100 * c + "%")
            },
            _updateCursorPosition: function(a) {
                var b = this._player.getDuration() ? a / this._player.getDuration() : 0;
                this.cursorNode && e.css(this.cursorNode, "left", 100 * b + "%")
            }
        });
        b.exports = i
    },
    {
        "../../lib/dom": 5,
        "../../lib/event": 6,
        "../../lib/function": 7,
        "../../lib/ua": 12,
        "../component": 17
    }],
    24 : [function(a, b, c) {
        var d = a("../component"),
        e = a("../../lib/util"),
        f = d.extend({
            init: function(a, b) {
                d.call(this, a, b),
                this.className = b.className ? b.className: "prism-time-display",
                this.addClass(this.className)
            },
            createEl: function() {
                var a = d.prototype.createEl.call(this, "div");
                return a.innerHTML = '<span class="current-time">00:00</span> <span class="time-bound">/</span> <span class="duration">00:00</span>',
                a
            },
            bindEvent: function() {
                var a = this;
                this._player.on("durationchange",
                function() {
                    var b = e.formatTime(a._player.getDuration());
                    b ? (document.querySelector("#" + a.id() + " .time-bound").style.display = "inline", document.querySelector("#" + a.id() + " .duration").style.display = "inline", document.querySelector("#" + a.id() + " .duration").innerText = b) : (document.querySelector("#" + a.id() + " .duration").style.display = "none", document.querySelector("#" + a.id() + " .time-bound").style.display = "none")
                }),
                this._player.on("timeupdate",
                function() {
                    var b = e.formatTime(a._player.getCurrentTime());
                    b ? (document.querySelector("#" + a.id() + " .current-time").style.display = "inline", document.querySelector("#" + a.id() + " .current-time").innerText = b) : document.querySelector("#" + a.id() + " .current-time").style.display = "none"
                })
            }
        });
        b.exports = f
    },
    {
        "../../lib/util": 14,
        "../component": 17
    }],
    25 : [function(a, b, c) {
        var d = a("../component"),
        e = d.extend({
            init: function(a, b) {
                d.call(this, a, b),
                this.addClass(b.className || "prism-volume")
            },
            bindEvent: function() {
                var a = this;
                this.on("click",
                function() {
                    a._player.muted() ? (a._player.unMute(), a.removeClass("mute")) : (a._player.mute(), a.addClass("mute"))
                })
            }
        });
        b.exports = e
    },
    {
        "../component": 17
    }],
    26 : [function(a, b, c) {
        b.exports = {
            bigPlayButton: a("./component/big-play-button"),
            controlBar: a("./component/controlbar"),
            progress: a("./component/progress"),
            playButton: a("./component/play-button"),
            liveDisplay: a("./component/live-display"),
            timeDisplay: a("./component/time-display"),
            fullScreenButton: a("./component/fullscreen-button"),
            volume: a("./component/volume")
        }
    },
    {
        "./component/big-play-button": 18,
        "./component/controlbar": 19,
        "./component/fullscreen-button": 20,
        "./component/live-display": 21,
        "./component/play-button": 22,
        "./component/progress": 23,
        "./component/time-display": 24,
        "./component/volume": 25
    }]
},
{},
[2]);