import requests
import time

cookies = {
    'sso_token': 'eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwczovL3VzZXIuemJsb2djbi5jb20vdXNlci9sb2dpbiIsImlhdCI6MTcyOTc2MDE3MSwiZXhwIjoxNzMyMzUyMTcxLCJuYmYiOjE3Mjk3NjAxNzEsImp0aSI6InhvaUpXRzFsSGZ6UTA4MFQxNzI5NzYwMTcxIiwic3ViIjoiNzUzM2FmOTQtYTM0Yi00ZTQ2LWE3ODAtMGMyMzcxZGE3MWFkIiwicHJ2IjoiNGFjMDE5ZWU2ZGIxZmM2MWZmMjZhMjMyNGU3YmY1OGE5MTg2NzY0MCIsInVzZXIiOnsiaWQiOiI3NTMzYWY5NC1hMzRiLTRlNDYtYTc4MC0wYzIzNzFkYTcxYWQiLCJ1c2VybmFtZSI6ImJiYmFhIiwibmlja25hbWUiOiJiYmJhYSIsImVtYWlsIjoiYmJiYWE2MzM4QGdtYWlsLmNvbSIsInZlcmlmaWVkIjoifGVtYWlsIiwicGVybWlzc2lvbnMiOiJ7fSIsInJvbGVzIjoiW10iLCJhdmF0YXJVcmwiOiJodHRwczovL2F2YXRhci56YmxvZ2NuLmNvbS9hdmF0YXIvaWQvNzUzM2FmOTQtYTM0Yi00ZTQ2LWE3ODAtMGMyMzcxZGE3MWFkIn19.LYjEx_omZRFiwpWisJGNfzMRoXv-JpWCIb2EIBUogfQ',
    'XSRF-TOKEN': 'eyJpdiI6ImZQRW00dGFoamVSNk5LTk82N2JtV1E9PSIsInZhbHVlIjoiS1R3K1pkcE1INTVuT0NORnF5ODY4QTg4NHZ3bFNKVERDRmF1UVh5K2lwRTlwRm5ROTdFNk81dmNtM0llaytnVXNObUpMVUIwUEllTFhBUWRneVpmQmErS2VKalBSdllkZzdLNGRZRlRBMFRDQ2s2eFlZUUxMVkNyb2RyVGhXbDMiLCJtYWMiOiIzMjE3Yjg0OThmMDhjMjlmMmQyYWMwOWQzNzRlOTFjM2VmZjFmYzI3Y2UxYTEyYjI1MzExYzRjNDA1ZWE0NjllIn0%3D',
    'uc_session': 'eyJpdiI6ImhMMGZXbVV4cUloSnc4VTRIcWUyblE9PSIsInZhbHVlIjoiUzREZWczdFdPbWhNSGF5dVdWTDBncnQ3VjBXZlJoVEp4dGVuMzFiR1EyUHNOc0xPSlg3UjR5ZWpzMUpHYjNkcTNYRCs0MDRhcTdFVFUvdk9vVlJzV3ZVK1JzdkQ1KzRFd2sxYU9BU1hHY0cxRWVTZ3VwQlFDaDJiTytYTjl2c0kiLCJtYWMiOiIxZTRkMTAwNGMwMGQ5ZjhkZjNlZGViODZhNDk3NjdmMzlmZDdjN2Y0ZDg0Y2FiY2Q0MTY0ZDFmYjQ4MDBjMTA2In0%3D',
}

headers = {
    'accept': 'application/json, text/plain, */*',
    'accept-language': 'zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6',
    'cache-control': 'no-cache',
    'content-type': 'application/json;charset=UTF-8',
    # 'cookie': 'sso_token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwczovL3VzZXIuemJsb2djbi5jb20vdXNlci9sb2dpbiIsImlhdCI6MTcyOTc2MDE3MSwiZXhwIjoxNzMyMzUyMTcxLCJuYmYiOjE3Mjk3NjAxNzEsImp0aSI6InhvaUpXRzFsSGZ6UTA4MFQxNzI5NzYwMTcxIiwic3ViIjoiNzUzM2FmOTQtYTM0Yi00ZTQ2LWE3ODAtMGMyMzcxZGE3MWFkIiwicHJ2IjoiNGFjMDE5ZWU2ZGIxZmM2MWZmMjZhMjMyNGU3YmY1OGE5MTg2NzY0MCIsInVzZXIiOnsiaWQiOiI3NTMzYWY5NC1hMzRiLTRlNDYtYTc4MC0wYzIzNzFkYTcxYWQiLCJ1c2VybmFtZSI6ImJiYmFhIiwibmlja25hbWUiOiJiYmJhYSIsImVtYWlsIjoiYmJiYWE2MzM4QGdtYWlsLmNvbSIsInZlcmlmaWVkIjoifGVtYWlsIiwicGVybWlzc2lvbnMiOiJ7fSIsInJvbGVzIjoiW10iLCJhdmF0YXJVcmwiOiJodHRwczovL2F2YXRhci56YmxvZ2NuLmNvbS9hdmF0YXIvaWQvNzUzM2FmOTQtYTM0Yi00ZTQ2LWE3ODAtMGMyMzcxZGE3MWFkIn19.LYjEx_omZRFiwpWisJGNfzMRoXv-JpWCIb2EIBUogfQ; XSRF-TOKEN=eyJpdiI6ImZQRW00dGFoamVSNk5LTk82N2JtV1E9PSIsInZhbHVlIjoiS1R3K1pkcE1INTVuT0NORnF5ODY4QTg4NHZ3bFNKVERDRmF1UVh5K2lwRTlwRm5ROTdFNk81dmNtM0llaytnVXNObUpMVUIwUEllTFhBUWRneVpmQmErS2VKalBSdllkZzdLNGRZRlRBMFRDQ2s2eFlZUUxMVkNyb2RyVGhXbDMiLCJtYWMiOiIzMjE3Yjg0OThmMDhjMjlmMmQyYWMwOWQzNzRlOTFjM2VmZjFmYzI3Y2UxYTEyYjI1MzExYzRjNDA1ZWE0NjllIn0%3D; uc_session=eyJpdiI6ImhMMGZXbVV4cUloSnc4VTRIcWUyblE9PSIsInZhbHVlIjoiUzREZWczdFdPbWhNSGF5dVdWTDBncnQ3VjBXZlJoVEp4dGVuMzFiR1EyUHNOc0xPSlg3UjR5ZWpzMUpHYjNkcTNYRCs0MDRhcTdFVFUvdk9vVlJzV3ZVK1JzdkQ1KzRFd2sxYU9BU1hHY0cxRWVTZ3VwQlFDaDJiTytYTjl2c0kiLCJtYWMiOiIxZTRkMTAwNGMwMGQ5ZjhkZjNlZGViODZhNDk3NjdmMzlmZDdjN2Y0ZDg0Y2FiY2Q0MTY0ZDFmYjQ4MDBjMTA2In0%3D',
    'origin': 'https://user.zblogcn.com',
    'pragma': 'no-cache',
    'priority': 'u=1, i',
    'referer': 'https://user.zblogcn.com/user/security/token',
    'sec-ch-ua': '"Chromium";v="130", "Microsoft Edge";v="130", "Not?A_Brand";v="99"',
    'sec-ch-ua-mobile': '?0',
    'sec-ch-ua-platform': '"Windows"',
    'sec-fetch-dest': 'empty',
    'sec-fetch-mode': 'cors',
    'sec-fetch-site': 'same-origin',
    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0',
    'x-csrf-token': 'TXzvruyVFyN0N4aQkeu5Tjw9USRrJZyr5TdfNEtl',
    'x-requested-with': 'XMLHttpRequest',
    'x-xsrf-token': 'eyJpdiI6ImZQRW00dGFoamVSNk5LTk82N2JtV1E9PSIsInZhbHVlIjoiS1R3K1pkcE1INTVuT0NORnF5ODY4QTg4NHZ3bFNKVERDRmF1UVh5K2lwRTlwRm5ROTdFNk81dmNtM0llaytnVXNObUpMVUIwUEllTFhBUWRneVpmQmErS2VKalBSdllkZzdLNGRZRlRBMFRDQ2s2eFlZUUxMVkNyb2RyVGhXbDMiLCJtYWMiOiIzMjE3Yjg0OThmMDhjMjlmMmQyYWMwOWQzNzRlOTFjM2VmZjFmYzI3Y2UxYTEyYjI1MzExYzRjNDA1ZWE0NjllIn0=',
}

json_data = {}
for i in range(1):
    time.sleep(1)
    response = requests.post('https://user.zblogcn.com/user/security/token', cookies=cookies, headers=headers, json=json_data)
    print("bbbaa|||" + response.text)