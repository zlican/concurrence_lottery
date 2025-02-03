<template>
    <div>
        <LuckyWheel
        ref="myLucky"
        :width="800"
        :height="800"
        :prizes="prizes"
        :blocks="blocks"
        :buttons="buttons"
        @start="startCallback"
        @end="endCallback"
    />
    </div>
  </template>
  
  <script>
  import { api } from '../axiosAPI/api'
  export default {
    data () {
      return {
        blocks: [{ padding: '13px', background: '#617df2' }],
        prizes: [],
        buttons: [{
          radius: '35%',
          background: '#8a9bf3',
          pointer: true,
          fonts: [{ text: '开始', top: '-10px' }]
        }],
        info: {0:0 ,10:1, 11:2, 12:3, 13:4, 15:5, 16:6, 17:7, 18:8, 19:9}
      }
    },
    methods: {
      // 点击抽奖按钮会触发star回调
      async startCallback () {
        // 调用抽奖组件的play方法开始游戏
        const res = await api.lottery()
        this.$refs.myLucky.play()
        const index = this.info[res.data.id]
        this.$refs.myLucky.stop(index)

      },
      // 抽奖结束会触发end回调
      endCallback (prize) {
        console.log(prize)
      },
      async getAllPrize(){
        const res = await api.getAllPrize()
        for(let i = 0;i<res.data.inventorys.length;i++){
            this.prizes.push({
              fonts:[{text:res.data.inventorys[i].Name,top:'10%'}],
              background: '#e9e8fe',
              imgs:[{src: 'src/assets/'+ res.data.inventorys[i].Picture,top:'25%', width: '100px', height: '100px' }]
            })
        }
      }
    },
    mounted(){
        this.getAllPrize()
    }
  }
  </script>

<style>

</style>