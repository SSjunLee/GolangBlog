import requests

base = "http://localhost:8080"


def TestApiGetPostInPage(cid, pid, psize):
    Url = base + "/api/post/page"
    res = requests.get(Url, params={
        'cid': cid,
        'pid': pid,
        'psize': psize
    })
    print(res.json())
    items = res.json()['data']
    for e in items:
        print(e)
    res.close()


def TestApiGetPageInPage(pid, psize):
    Url = base + "/api/page/page"
    res = requests.get(Url, params={
        'pid': pid,
        'psize': psize
    })
    print(res.json())
    items = res.json()['data']
    for e in items:
        print(e)
    res.close()


def TestLogin():
    Url = base + "/api/login"
    res = requests.post(Url, {"pp": "qq"})
    print(res.json())


if __name__ == '__main__':
    # TestApiGetPostInPage(2, 1, 5)
    # TestApiGetPageInPage(1,5)
    TestLogin()
