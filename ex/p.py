import requests

headers = {
    'Accept': 'application/json, text/plain, */*',
    'Accept-Language': 'zh',
    'Connection': 'keep-alive',
    'Content-Type': 'application/json',
    'EntranceCode': 'anNq',
    'Origin': 'http://38.181.29.193:34567',
    'Referer': 'http://38.181.29.193:34567/jsj',
    'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36',
}

json_data = {
    'name': 'wiki888',
    'password': 'L5Zz4687UfFhKcH7',
    'ignoreCaptcha': True,
    'captcha': '',
    'captchaID': 'CNOMr9mDfqNhJjoHGbQX',
    'authMethod': 'session',
    'language': 'zh',
}

response = requests.post('http://38.181.29.193:34567/api/v1/auth/login', headers=headers, json=json_data, verify=False)
print(response.text)
# Note: json_data will not be serialized by requests
# exactly as it was in the original request.
#data = '{"name":"wiki888","password":"L5Zz4687UfFhKcH7","ignoreCaptcha":true,"captcha":"","captchaID":"CNOMr9mDfqNhJjoHGbQX","authMethod":"session","language":"zh"}'
#response = requests.post('http://38.181.29.193:34567/api/v1/auth/login', headers=headers, data=data, verify=False)
