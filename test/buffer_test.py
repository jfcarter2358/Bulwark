import pytest
import requests
from config import *
from requests.packages.urllib3.exceptions import InsecureRequestWarning
requests.packages.urllib3.disable_warnings(InsecureRequestWarning)

def test_buffer_create():
    name = "foobar1"
    headers = {"X-Bulwark-API" : f'{BULWARK_SECRET_KEY}' }

    response = requests.post(f"{BULWARK_PROTOCOL}://{BULWARK_HOST}:{BULWARK_PORT}/api/v1/buffer/{name}", headers=headers, verify=False)
    assert response.status_code < 400

    response = requests.delete(f"{BULWARK_PROTOCOL}://{BULWARK_HOST}:{BULWARK_PORT}/api/v1/buffer/{name}", headers=headers, verify=False)
    assert response.status_code < 400

def test_buffer_set():
    name = "foobar2"
    data = {"data": "foobar1"}
    headers = {"X-Bulwark-API" : f'{BULWARK_SECRET_KEY}' }

    response = requests.post(f"{BULWARK_PROTOCOL}://{BULWARK_HOST}:{BULWARK_PORT}/api/v1/buffer/{name}", headers=headers, verify=False)
    assert response.status_code < 400

    response = requests.put(f"{BULWARK_PROTOCOL}://{BULWARK_HOST}:{BULWARK_PORT}/api/v1/buffer/{name}", headers=headers, json=data, verify=False)
    assert response.status_code < 400

    response = requests.delete(f"{BULWARK_PROTOCOL}://{BULWARK_HOST}:{BULWARK_PORT}/api/v1/buffer/{name}", headers=headers, verify=False)
    assert response.status_code < 400

def test_buffer_get():
    name = "foobar2"
    data1 = {"data": "foobar1"}
    data2 = {"data": "foobar2"}
    headers = {"X-Bulwark-API" : f'{BULWARK_SECRET_KEY}' }

    response = requests.post(f"{BULWARK_PROTOCOL}://{BULWARK_HOST}:{BULWARK_PORT}/api/v1/buffer/{name}", headers=headers, verify=False)
    assert response.status_code < 400

    response = requests.put(f"{BULWARK_PROTOCOL}://{BULWARK_HOST}:{BULWARK_PORT}/api/v1/buffer/{name}", headers=headers, json=data1, verify=False)
    assert response.status_code < 400

    response = requests.put(f"{BULWARK_PROTOCOL}://{BULWARK_HOST}:{BULWARK_PORT}/api/v1/buffer/{name}", headers=headers, json=data2, verify=False)
    assert response.status_code < 400

    response = requests.get(f"{BULWARK_PROTOCOL}://{BULWARK_HOST}:{BULWARK_PORT}/api/v1/buffer/{name}", headers=headers, verify=False)
    assert response.json()["data"] == data2["data"]

    response = requests.delete(f"{BULWARK_PROTOCOL}://{BULWARK_HOST}:{BULWARK_PORT}/api/v1/buffer/{name}", headers=headers, verify=False)
    assert response.status_code < 400

def test_buffer_delete():
    name = "foobar1"
    headers = {"X-Bulwark-API" : f'{BULWARK_SECRET_KEY}' }

    response = requests.post(f"{BULWARK_PROTOCOL}://{BULWARK_HOST}:{BULWARK_PORT}/api/v1/buffer/{name}", headers=headers, verify=False)
    assert response.status_code < 400

    response = requests.delete(f"{BULWARK_PROTOCOL}://{BULWARK_HOST}:{BULWARK_PORT}/api/v1/buffer/{name}", headers=headers, verify=False)
    assert response.status_code < 400
