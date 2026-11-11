# Opt-in graphical fixture runner. Only its two disposable windows are changed.
# Run from any directory; GO_BINARY may point to a local Go toolchain.
import os,sys,json,subprocess,time
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
if len(sys.argv)>1:
 import gi
 gi.require_version('Gtk','3.0')
 from gi.repository import Gtk,Gio
 app=Gtk.Application(application_id='org.codex.HyprDockProbe',flags=Gio.ApplicationFlags.NON_UNIQUE)
 def activate(app):
  win=Gtk.ApplicationWindow(application=app,title='Teste temporário hypr-dock')
  win.set_default_size(400,240)
  area=Gtk.DrawingArea()
  def draw(w,c):
   c.set_source_rgb(*( (1,0,0) if sys.argv[1]=='red' else (0,0,1) ));c.paint()
  area.connect('draw',draw);win.add(area);win.show_all()
 app.connect('activate',activate);app.run([]);sys.exit()
def q(k):return json.loads(subprocess.check_output(['hyprctl','-j',k]))
def d(x):
 o=subprocess.check_output(['hyprctl','dispatch',x],text=True).strip();assert o=='ok',o
original=q('activewindow').get('address');cursor=q('cursorpos');children=[];addresses=[]
monitor=next(m for m in q('monitors') if m['focused'])
if monitor['scale'] != 1 or monitor['transform'] != 0:
 raise SystemExit('This fixture runner requires an unrotated scale-1 monitor.')
if monitor['width'] < 640 or monitor['height'] < 480:
 raise SystemExit('The monitor is too small for the overlapping fixture windows.')
x=monitor['x']+(monitor['width']-400)//2
y=monitor['y']+(monitor['height']-240)//2
try:
 for color in ['red','blue']:
  p=subprocess.Popen([sys.executable,__file__,color],stdout=subprocess.DEVNULL,stderr=subprocess.DEVNULL);children.append(p)
  for _ in range(60):
   w=next((w for w in q('clients') if w['pid']==p.pid),None)
   if w:break
   time.sleep(.05)
  assert w
  print('Fixture class:',w['class'],flush=True)
  a=w['address'];addresses.append(a)
  if not w['floating']:d(f'hl.dsp.window.float({{window="address:{a}",action="on"}})')
  d(f'hl.dsp.window.resize({{window="address:{a}",x=400,y=240,relative=false}})')
  d(f'hl.dsp.window.move({{window="address:{a}",x={x},y={y},relative=false}})')
  time.sleep(.25)
 d(f'hl.dsp.cursor.move({{x={monitor["x"]+monitor["width"]//2},y={monitor["y"]+monitor["height"]-2}}})')
 env=os.environ.copy();env.update(HYPR_DOCK_TEST_TARGET=addresses[0])
 subprocess.run([os.environ.get('GO_BINARY','go'),'test','./pkg/ipc','-run','TestLuaFocusIntegration','-count=1','-v'],cwd=ROOT,env=env,check=True)
finally:
 for p in children:p.terminate()
 for p in children:p.wait(timeout=4)
 if original:
  d('(function() local old=hl.get_config("cursor.no_warps");hl.config({cursor={no_warps=true}});local ok,err=pcall(hl.dispatch,hl.dsp.focus({window="address:'+original+'"}));hl.config({cursor={no_warps=old}});if not ok then error(err) end end)')

 d(f'hl.dsp.cursor.move({{x={cursor["x"]},y={cursor["y"]}}})')
