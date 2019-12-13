from aiohttp import log
import asyncio
import websockets
import rapidjson as json
import traceback

class ConnectionClosed(Exception):
    pass

class DpowClient(object):
    """Websocket client for DPoW & BoomPoW"""
    DPOW_SERVER = 'wss://dpow.nanocenter.org/service_ws/'
    BPOW_SERVER = 'wss://bpow.banano.cc/service_ws/'
    NANO_DIFFICULTY_CONST = 'ffffffc000000000'

    def __init__(self, dpow_user: str, dpow_key: str, force_nano_difficulty: bool = False, work_futures: dict = {}, bpow: bool = False):
        """work_futures is a dict of id, futures to keep track of requests and responses"""
        self.bpow = bpow
        self.url = self.BPOW_SERVER if self.bpow else self.DPOW_SERVER
        self.ws: websockets.WebSocketClientProtocol = None
        self.stop = False
        self.dpow_user = dpow_user
        self.dpow_key = dpow_key
        self.work_futures = work_futures
        self.difficulty = self.NANO_DIFFICULTY_CONST if force_nano_difficulty else None

    async def request_work(self, id: str, hash: str, difficulty: str = None):
        """Request work from DPoW/BPoW WS"""
        if self.stop or self.ws is None:
            return
        try:
            if self.ws.closed:
                raise ConnectionClosed()
            req = {
                "user": self.dpow_user,
                "api_key": self.dpow_key,
                "hash": hash,
                "id": id,
                "timeout": 15
            }
            if difficulty is not None:
                req['difficulty'] = difficulty
            elif self.difficulty is not None:
                req['difficulty'] = self.difficulty
            await self.ws.send(json.dumps(req))
        except Exception:
            raise ConnectionClosed()        

    async def setup(self):
        try:
            self.ws = await websockets.connect(self.url)
            log.server_logger.info(f"Connected to {'BoomPoW' if self.bpow else 'Distributed PoW'} Service")
        except Exception as e:
            log.server_logger.critical("Error connecting to websocket server.")
            log.server_logger.error(traceback.format_exc())
            raise

    async def close(self):
        self.stop = True
        await self.ws.wait_closed()

    async def reconnect_forever(self):
        log.server_logger.warn("Attempting websocket reconnection every 30 seconds...")
        while not self.stop:
            try:
                await self.setup()
                log.server_logger.warn("Connected to websocket!")
                break
            except:
                log.server_logger.debug("Websocket reconnection failed")
                await asyncio.sleep(30)

    async def loop(self):
        while not self.stop:
            try:
                rec = json.loads(await self.ws.recv())
                request_id = rec.get("id", None)
                if request_id is not None:
                    try:
                        result = self.work_futures[str(request_id)]
                        if not result.done():
                            result.set_result(rec)
                    except KeyError:
                        pass                    
            except KeyboardInterrupt:
                break
            except websockets.exceptions.ConnectionClosed as e:
                log.server_logger.error(f"Connection closed to websocket. Code: {e.code} , reason: {e.reason}.")
                await self.reconnect_forever()
            except Exception as e:
                log.server_logger.critical(f"Unknown exception while handling getting a websocket message:\n{traceback.format_exc()}")
                await self.reconnect_forever()